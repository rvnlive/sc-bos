package pubcache

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/smart-core-os/sc-api/go/traits"
	"github.com/smart-core-os/sc-golang/pkg/resource"
	"github.com/smart-core-os/sc-golang/pkg/trait/publication"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestCache_Pull(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	id := "foo"
	initial := &traits.Publication{
		Id:          id,
		Version:     "1",
		Body:        []byte(`{"hello": "world"}`),
		PublishTime: timestamppb.Now(),
		MediaType:   "application/json",
	}
	model := publication.NewModel(
		resource.WithInitialRecord(id, initial),
	)
	modelServer := publication.NewModelServer(model)
	modelClient := publication.WrapApi(modelServer)

	cache := New(ctx, modelClient, "", id, nil)

	ch := cache.Pull(ctx, id)

	// first value received should be the initial value
	val := recvTimeout(ch, time.Second)
	if !proto.Equal(initial, val) {
		t.Errorf("initial value mismatch (-want +got):\n%s", cmp.Diff(initial, val))
	}

	// change the publication in the model
	updated := &traits.Publication{
		Id:          id,
		Version:     "2",
		Body:        []byte(`{"foo":"bar"}`),
		PublishTime: timestamppb.Now(),
		MediaType:   "application/json",
	}
	go func() {
		_, err := model.UpdatePublication(id, updated)
		if err != nil {
			panic(err)
		}
	}()

	// check the new value has been propagated
	val = recvTimeout(ch, time.Second)
	if !proto.Equal(updated, val) {
		t.Errorf("updated value mismatch (-want +got):\n%s", cmp.Diff(updated, val))
	}
}

func recvTimeout[T any](ch <-chan T, timeout time.Duration) T {
	select {
	case <-time.After(timeout):
		panic("receive timeout")
	case val := <-ch:
		return val
	}
}
