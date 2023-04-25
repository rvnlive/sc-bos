package statuspb

import (
	"context"
	"sort"
	"strings"
	"time"

	"golang.org/x/exp/slices"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/smart-core-os/sc-golang/pkg/cmp"
	"github.com/smart-core-os/sc-golang/pkg/masks"
	"github.com/smart-core-os/sc-golang/pkg/resource"
	"github.com/vanti-dev/sc-bos/pkg/gen"
	"github.com/vanti-dev/sc-bos/pkg/util/chans"
)

// Model provides an in-memory model for storing and retrieving problems as a status log.
// A Model can be used as a backing store for the StatusApi.
type Model struct {
	problems *resource.Collection // of *gen.StatusLog_Problem

	equivalence cmp.Message
}

// NewModel returns a new Model.
func NewModel(opts ...resource.Option) *Model {
	return &Model{
		problems:    resource.NewCollection(opts...),
		equivalence: cmp.Equal(),
	}
}

// GetCurrentStatus returns all known problems as a status log.
func (m *Model) GetCurrentStatus(readMask *fieldmaskpb.FieldMask) (*gen.StatusLog, error) {
	out := m.problemsToStatusLog(m.problems.List())

	filter := masks.NewResponseFilter(masks.WithFieldMask(readMask))
	clone := filter.FilterClone(out)
	return clone.(*gen.StatusLog), nil
}

func (m *Model) problemsToStatusLog(problemMsgs []proto.Message) *gen.StatusLog {
	pm := &ProblemMerger{}
	pm.AddProblemMessages(problemMsgs)
	return pm.Build()
}

// UpdateProblem will add or update the given problem in the model.
// Pull methods will be notified.
func (m *Model) UpdateProblem(problem *gen.StatusLog_Problem) (*gen.StatusLog_Problem, error) {
	if problem.RecordTime == nil {
		problem.RecordTime = timestamppb.New(m.problems.Clock().Now())
	}
	res, err := m.problems.Update(problem.Name, problem, resource.WithCreateIfAbsent(), resource.InterceptAfter(func(old, new proto.Message) {
		if old == nil {
			return
		}
		var oldp, newp *gen.StatusLog_Problem
		oldp = old.(*gen.StatusLog_Problem)
		newp = new.(*gen.StatusLog_Problem)
		if oldp.RecordTime == nil {
			return
		}
		if oldp.Level == newp.Level && oldp.RecordTime.AsTime().Before(newp.RecordTime.AsTime()) {
			newp.RecordTime = oldp.RecordTime
		}
	}))
	if err != nil {
		return nil, err
	}
	return res.(*gen.StatusLog_Problem), nil
}

// DeleteProblem removes the named problem if it exists.
func (m *Model) DeleteProblem(name string) {
	_, _ = m.problems.Delete(name, resource.WithAllowMissing(true))
}

type StatusLogChange struct {
	StatusLog  *gen.StatusLog
	ChangeTime time.Time
}

func (m *Model) PullCurrentStatus(ctx context.Context, readMask *fieldmaskpb.FieldMask, updatesOnly bool) <-chan StatusLogChange {
	// todo: convert the func arguments to resource.ReadOption when enough resource apis are published

	send := make(chan StatusLogChange)
	stream := m.problems.Pull(ctx)
	go func() {
		defer close(send)

		var lastSend *gen.StatusLog
		var problems []proto.Message // sorted by name
		filter := masks.NewResponseFilter(masks.WithFieldMask(readMask))
		seeding := true
		for change := range stream {
			i, found := sort.Find(len(problems), func(i int) int {
				return strings.Compare(problems[i].(*gen.StatusLog_Problem).Name, change.Id)
			})
			switch {
			case change.NewValue == nil:
				if found {
					problems = slices.Delete(problems, i, 1)
				}
			default:
				if found {
					problems[i] = change.NewValue
				} else {
					problems = slices.Insert(problems, i, change.NewValue)
				}
			}

			if change.LastSeedValue {
				seeding = false
			}
			if seeding || (updatesOnly && change.LastSeedValue) {
				continue
			}

			statusLog := m.problemsToStatusLog(problems)
			statusLog = filter.FilterClone(statusLog).(*gen.StatusLog)
			if m.equivalence(statusLog, lastSend) {
				continue
			}
			lastSend = statusLog
			err := chans.SendContext(ctx, send, StatusLogChange{
				StatusLog:  statusLog,
				ChangeTime: change.ChangeTime,
			})
			if err != nil {
				return
			}
		}
	}()
	return send
}
