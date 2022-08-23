package policy

import (
	"context"
	"embed"
	"io/fs"
	"strings"
	"sync"
	"time"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
)

type static struct {
	compiler *ast.Compiler
}

func (p *static) EvalPolicy(ctx context.Context, query string, input Attributes) (rego.ResultSet, error) {
	return rego.New(
		rego.Compiler(p.compiler),
		rego.Input(input),
		rego.Query(query),
	).Eval(ctx)
}

type cachedStatic struct {
	compiler *ast.Compiler
	cache    map[string]*regoCacheEntry
	cacheM   sync.Mutex
}

func newCachedStatic(compiler *ast.Compiler) *cachedStatic {
	return &cachedStatic{
		cache:    make(map[string]*regoCacheEntry),
		compiler: compiler,
	}
}

func (p *cachedStatic) EvalPolicy(ctx context.Context, query string, input Attributes) (rego.ResultSet, error) {
	partial, err := p.loadPartialCached(ctx, query)
	if err != nil {
		return nil, err
	}

	return partial.Rego(rego.Input(input)).Eval(ctx)
}

// Results are cached; the partial evaluation is performed once the first time and re-used for subsequent calls.
// If the provided context is cancelled before the result is ready, the process will continue in the background and
// the context error is returned.
// If loadPartialCached returns a non-context error, then future calls with the same query will always return the same error.
func (p *cachedStatic) loadPartialCached(ctx context.Context, query string) (rego.PartialResult, error) {
	p.cacheM.Lock()
	entry, ok := p.cache[query]
	if !ok {
		entry = &regoCacheEntry{done: make(chan struct{})}
		p.cache[query] = entry
	}
	p.cacheM.Unlock()

	// each cache entry only gets one change to compile - it's a deterministic process, so if it fails once there's no
	// point trying again later
	entry.once.Do(func() {
		// run asynchronously so the compilation can complete in the background if ctx is cancelled early
		go func() {
			defer close(entry.done)
			// if a policy file takes more than 5 seconds to compile, something is wrong with it
			bgctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			r := rego.New(
				rego.Compiler(p.compiler),
				rego.Query(query),
			)
			entry.partialResult, entry.err = r.PartialResult(bgctx)
		}()
	})
	select {
	case <-entry.done:
		return entry.partialResult, entry.err
	case <-ctx.Done():
		return rego.PartialResult{}, ctx.Err()
	}
}

type regoCacheEntry struct {
	once          sync.Once
	done          chan struct{}
	partialResult rego.PartialResult
	err           error
}

func compileFS(sources fs.FS) (*ast.Compiler, error) {
	files := make(map[string]string)
	err := fs.WalkDir(sources, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !strings.HasSuffix(path, ".rego") {
			return nil
		}

		contents, err := fs.ReadFile(sources, path)
		if err != nil {
			return err
		}

		files[path] = string(contents)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return ast.CompileModules(files)
}

var (
	//go:embed default
	defaultPolicyFS embed.FS
	defaultCompiler *ast.Compiler
)

func init() {
	compiler, err := compileFS(defaultPolicyFS)
	if err != nil {
		panic(err)
	}
	defaultCompiler = compiler
}

func Default() Policy {
	return newCachedStatic(defaultCompiler)
}

func FromFS(f fs.FS) (Policy, error) {
	compiler, err := compileFS(f)
	if err != nil {
		return nil, err
	}

	return newCachedStatic(compiler), nil
}
