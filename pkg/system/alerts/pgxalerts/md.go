package pgxalerts

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/smart-core-os/sc-golang/pkg/resource"
	"github.com/vanti-dev/sc-bos/pkg/gen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) GetAlertMetadata(ctx context.Context, request *gen.GetAlertMetadataRequest) (*gen.AlertMetadata, error) {
	if err := s.initAlertMetadata(ctx); err != nil {
		return nil, err
	}
	return s.md.Get(resource.WithReadMask(request.ReadMask)).(*gen.AlertMetadata), nil
}

func (s *Server) PullAlertMetadata(request *gen.PullAlertMetadataRequest, server gen.AlertApi_PullAlertMetadataServer) error {
	if err := s.initAlertMetadata(server.Context()); err != nil {
		return err
	}
	for change := range s.md.Pull(server.Context(), resource.WithReadMask(request.GetReadMask()), resource.WithUpdatesOnly(request.UpdatesOnly)) {
		err := server.Send(&gen.PullAlertMetadataResponse{Changes: []*gen.PullAlertMetadataResponse_Change{
			{
				Name:       request.Name,
				ChangeTime: timestamppb.New(change.ChangeTime),
				Metadata:   change.Value.(*gen.AlertMetadata),
			},
		}})
		if err != nil {
			return err
		}
	}

	return nil
}

// initAlertMetadata seeds s.md and keeps it up to date with changes to the underlying store
// if s.md hasn't already been seeded.
// initAlertMetadata blocks until either s.md is seeded or an error occurs or ctx expires.
// If an error occurs seeding s.md then it is returned
func (s *Server) initAlertMetadata(ctx context.Context) error {
	s.mdMu.Lock()
	defer s.mdMu.Unlock()
	wait := func() error {
		select {
		case <-s.mdC:
			return s.mdErr
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	if s.mdC != nil {
		// init is either in progress or has completed
		return wait()
	}

	// need to init the alert md
	s.mdC = make(chan struct{})
	ctx, s.mdStop = context.WithCancel(context.Background())
	go func() {
		// no need to lock here as we already know we're the only ones writing to mdC or mdErr
		defer close(s.mdC)

		// These are populated in the db transaction to make sure we don't miss any events.
		// There is a race condition here, though very unlikely (I think).
		// There is space between our transaction and other RPC transactions (like AckAlert)
		// for events to be missed or duplicated.
		// I apologise to any future maintainer who find this comment and has to fix it, likely under pressure.
		var events <-chan *gen.PullAlertsResponse_Change
		var eventsCtx context.Context
		var eventsCancel context.CancelFunc

		// Collect initial stats from the DB
		md := &gen.AlertMetadata{
			AcknowledgedCounts: make(map[bool]uint32),
			FloorCounts:        make(map[string]uint32),
			ZoneCounts:         make(map[string]uint32),
			SeverityCounts:     make(map[int32]uint32),
		}
		err := s.pool.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
			// totals
			err := tx.QueryRow(ctx, `SELECT COUNT(*) FROM alerts`).Scan(&md.TotalCount)
			if err != nil {
				return fmt.Errorf("count all %w", err)
			}

			// ack and nak counts
			var ackCount uint32
			err = tx.QueryRow(ctx, `SELECT COUNT(*) FROM alerts WHERE ack_time IS NOT NULL`).Scan(&ackCount)
			if err != nil {
				return fmt.Errorf("nak count %w", err)
			}
			md.AcknowledgedCounts = map[bool]uint32{
				true:  ackCount,
				false: md.TotalCount - ackCount,
			}

			// floors
			floors, err := queryGroupCounts(ctx, tx, "floor", s.Floors)
			if err != nil {
				return fmt.Errorf("floors %w", err)
			}
			md.FloorCounts = floors
			// zones
			zones, err := queryGroupCounts(ctx, tx, "zone", s.Floors)
			if err != nil {
				return fmt.Errorf("zones %w", err)
			}
			md.ZoneCounts = zones
			// severity
			severity, err := queryGroupCounts(ctx, tx, "severity", s.Severity)
			if err != nil {
				return fmt.Errorf("severity %w", err)
			}
			md.SeverityCounts = severity

			if eventsCancel != nil {
				eventsCancel()
			}
			eventsCtx, eventsCancel = context.WithCancel(ctx)
			events = s.bus.Listen(eventsCtx)

			return nil
		})
		if err != nil {
			s.mdErr = err
			s.mdC = nil // allow another call to attempt again
			return
		}
		s.md = resource.NewValue(resource.WithInitialValue(md))

		// setup listeners for changes to the DB so we can track those changes in metadata
		go func() {
			defer eventsCancel()
			for {
				select {
				case <-ctx.Done():
					return // the server is closing
				case e := <-events:
					applyMdDelta(s.md, e)
				}
			}
		}()
	}()
	return wait()
}

func applyMdDelta(md *resource.Value, e *gen.PullAlertsResponse_Change) error {
	// note: AlertMetadata uses uint32 so we can't use it to store our deltas
	_, err := md.Set(&gen.AlertMetadata{}, resource.InterceptBefore(func(old, new proto.Message) {
		oldMd, newMd := old.(*gen.AlertMetadata), new.(*gen.AlertMetadata)
		proto.Merge(newMd, oldMd)

		// total
		if e.OldValue == nil && e.NewValue != nil {
			newMd.TotalCount++
		} else if e.OldValue != nil && e.NewValue == nil {
			newMd.TotalCount--
			if newMd.TotalCount < 0 { // shouldn't be needed, but just in case
				newMd.TotalCount = 0
			}
		}

		// ack/nak
		if e.OldValue == nil && e.NewValue != nil {
			acked := e.NewValue.GetAcknowledgement().GetAcknowledgeTime() != nil
			if acked {
				newMd.AcknowledgedCounts[true]++
			} else {
				newMd.AcknowledgedCounts[false]++
			}
		} else if e.OldValue != nil && e.NewValue == nil {
			acked := e.OldValue.GetAcknowledgement().GetAcknowledgeTime() != nil
			if acked {
				newMd.AcknowledgedCounts[true]--
			} else {
				newMd.AcknowledgedCounts[false]--
			}
		} else {
			oldAck, newAck := e.GetOldValue().GetAcknowledgement().GetAcknowledgeTime(),
				e.GetNewValue().GetAcknowledgement().GetAcknowledgeTime()
			if oldAck == nil && newAck != nil {
				newMd.AcknowledgedCounts[true]++
				if newMd.AcknowledgedCounts[false] > 0 { // just in case
					newMd.AcknowledgedCounts[false]--
				}
			} else if oldAck != nil && newAck == nil {
				if newMd.AcknowledgedCounts[true] > 0 { // just in case
					newMd.AcknowledgedCounts[true]--
				}
				newMd.AcknowledgedCounts[false]++
			}
		}

		// floors, zones, and severity
		mapDelta(e.GetOldValue().GetFloor(), e.GetNewValue().GetFloor(), newMd.FloorCounts)
		mapDelta(e.GetOldValue().GetZone(), e.GetNewValue().GetZone(), newMd.ZoneCounts)
		mapDelta(int32(e.GetOldValue().GetSeverity()), int32(e.GetNewValue().GetSeverity()), newMd.SeverityCounts)
	}))

	return err
}

func mapDelta[T comparable](o, n T, m map[T]uint32) {
	if o == n {
		return
	}
	var zero T
	if o != zero {
		if c, ok := m[o]; ok && c > 0 {
			m[o] = c - 1
		}
	}
	if n != zero {
		m[n]++
	}
}

func queryGroupCounts[K comparable](ctx context.Context, tx pgx.Tx, col string, seed []K) (map[K]uint32, error) {
	all := make(map[K]uint32)
	for _, floor := range seed {
		all[floor] = 0
	}
	query, err := tx.Query(ctx, fmt.Sprintf(`SELECT %s, COUNT(id) FROM alerts GROUP BY %s`, col, col))
	if err != nil {
		return nil, fmt.Errorf("query %w", err)
	}
	for query.Next() {
		var id K
		var count uint32
		err = query.Scan(&id, &count)
		if err != nil {
			return nil, fmt.Errorf("row scan %w", err)
		}
		all[id] = count
	}
	return all, nil
}
