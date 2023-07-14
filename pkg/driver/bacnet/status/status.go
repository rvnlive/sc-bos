package status

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"go.uber.org/zap"
	"golang.org/x/exp/constraints"

	"github.com/vanti-dev/gobacnet"
	"github.com/vanti-dev/gobacnet/enum/eventstate"
	"github.com/vanti-dev/gobacnet/enum/reliability"
	"github.com/vanti-dev/gobacnet/property"
	"github.com/vanti-dev/gobacnet/types/objecttype"
	"github.com/vanti-dev/sc-bos/pkg/driver/bacnet/adapt"
	"github.com/vanti-dev/sc-bos/pkg/driver/bacnet/comm"
	"github.com/vanti-dev/sc-bos/pkg/driver/bacnet/config"
	"github.com/vanti-dev/sc-bos/pkg/driver/bacnet/known"
	"github.com/vanti-dev/sc-bos/pkg/gen"
	"github.com/vanti-dev/sc-bos/pkg/gentrait/statuspb"
)

type Monitor struct {
	client  *gobacnet.Client
	known   known.Context
	status  *statuspb.Map
	objects []*Object

	Logger *zap.Logger
}

func NewMonitor(client *gobacnet.Client, known known.Context, status *statuspb.Map) *Monitor {
	return &Monitor{
		Logger: zap.NewNop(),
		client: client,
		known:  known,
		status: status,
	}
}

// AddDevice adds all objects of the given device to the monitor.
// This must be called before starting the monitor.
func (m *Monitor) AddDevice(name string, device config.Device) {
	for _, obj := range device.Objects {
		m.AddDeviceObject(name, device, obj)
	}
}

// AddDeviceObject adds the given object to the monitor.
// This must be called before starting the monitor.
func (m *Monitor) AddDeviceObject(name string, device config.Device, obj config.Object) {
	dRef := config.NewDeviceRefID(device.ID)
	oRef := config.NewObjectRefID(obj.ID)
	on := adapt.ObjectName(obj)
	so := &Object{
		Name:         name,
		Test:         on,
		EventState:   &config.ValueSource{Property: propPtr(property.EventState), Device: dRef, Object: oRef},
		Reliability:  &config.ValueSource{Property: propPtr(property.Reliability), Device: dRef, Object: oRef},
		OutOfService: &config.ValueSource{Property: propPtr(property.OutOfService), Device: dRef, Object: oRef},
	}
	if canHaveLimits(obj) {
		so.LowLimit = &config.ValueSource{Property: propPtr(property.LowLimit), Device: dRef, Object: oRef}
		so.HighLimit = &config.ValueSource{Property: propPtr(property.HighLimit), Device: dRef, Object: oRef}
	}
	m.AddObject(so)
}

func canHaveLimits(obj config.Object) bool {
	// These values were collected by searching the BACnet spec for any mention of Low_Limit
	// and checking the "Properties of <some bacnet object> Type" tables.
	switch obj.ID.Type {
	case objecttype.AnalogInput, objecttype.AnalogOutput, objecttype.AnalogValue:
		return true
	case objecttype.LargeAnalogValue, objecttype.IntegerValue, objecttype.PositiveIntegerValue:
		return true
	case objecttype.PulseConverter, objecttype.Accumulator:
		return true
	default:
		return false
	}
}

// AddObject adds the given object to the monitor.
// This must be called before starting the monitor.
func (m *Monitor) AddObject(obj *Object) {
	m.objects = append(m.objects, obj)
}

func propPtr(p property.ID) *config.PropertyID {
	return (*config.PropertyID)(&p)
}

type Object struct {
	Name  string              `json:"name,omitempty"` // the name we associate problems with
	Test  string              `json:"test,omitempty"` // problem name suffix
	Level gen.StatusLog_Level `json:"level,omitempty"`

	// these help us to determine the health of the object
	EventState   *config.ValueSource `json:"eventState,omitempty"`
	Reliability  *config.ValueSource `json:"reliability,omitempty"`
	OutOfService *config.ValueSource `json:"outOfService,omitempty"`
	// these help with improving the quality of the status reporting for the above
	LowLimit  *config.ValueSource `json:"lowLimit,omitempty"`
	HighLimit *config.ValueSource `json:"highLimit,omitempty"`
	// These allow an object to explicitly represent a fault via its PresentValue.
	// If the value isn't equal to the nominalValue then we'll report a fault.
	Value        *config.ValueSource `json:"value,omitempty"`
	NominalValue any                 `json:"nominalValue,omitempty"`
}

func (o Object) hasLimits() bool {
	return o.LowLimit != nil && o.HighLimit != nil
}

func (o Object) hasValue() bool {
	return o.Value != nil
}

func (o Object) hasStatus() bool {
	return o.EventState != nil && o.Reliability != nil && o.OutOfService != nil
}

func (m *Monitor) Poll(ctx context.Context) error {
	allFields := make([]config.ValueSource, 0, len(m.objects)*3)
	for _, o := range m.objects {
		if o.hasStatus() {
			allFields = append(allFields, *o.EventState, *o.Reliability, *o.OutOfService)
		}
		if o.hasLimits() {
			allFields = append(allFields, *o.LowLimit, *o.HighLimit)
		}
		if o.hasValue() {
			allFields = append(allFields, *o.Value)
		}
	}

	var resultsCursor int
	results := comm.ReadPropertiesChunked(ctx, m.client, m.known, 30, allFields...)
	for _, o := range m.objects {
		var reqs []string  // requests we check
		notFoundCount := 0 // total responses that indicate the property doesn't exist
		timeoutCount := 0  // total responses that indicate the property timed out

		var errs []error
		handleErr := func(prop string, err error) {
			reqs = append(reqs, prop)
			switch {
			case errors.Is(err, comm.ErrPropNotFound):
				notFoundCount++
			case errors.Is(err, context.DeadlineExceeded):
				timeoutCount++
			case err != nil:
				errs = append(errs, comm.ErrReadProperty{Cause: err, Prop: prop})
			}
		}

		var err error
		var (
			eventStateVal   eventstate.EventState
			reliabilityVal  reliability.Reliability
			outOfServiceVal bool
		)
		hasStatus := o.hasStatus()
		if hasStatus {
			eventStateVal, err = asEnumType[eventstate.EventState](results[resultsCursor])
			resultsCursor++
			handleErr("EventState", err)

			reliabilityVal, err = asEnumType[reliability.Reliability](results[resultsCursor])
			resultsCursor++
			handleErr("Reliability", err)

			outOfServiceVal, err = comm.BoolValue(results[resultsCursor])
			resultsCursor++
			handleErr("OutOfService", err)
		}

		var lowLimitVal, highLimitVal float32
		hasLimits := o.hasLimits()
		if hasLimits {
			lowLimitVal, err = comm.Float32Value(results[resultsCursor])
			resultsCursor++
			handleErr("LowLimit", err)
			if errors.Is(err, comm.ErrPropNotFound) {
				hasLimits = false
			}

			highLimitVal, err = comm.Float32Value(results[resultsCursor])
			resultsCursor++
			handleErr("HighLimit", err)
			if errors.Is(err, comm.ErrPropNotFound) {
				hasLimits = false
			}
		}

		var valueVal any
		hasValue := o.hasValue()
		if hasValue {
			valueVal, err = comm.AnyValue(results[resultsCursor])
			resultsCursor++
			handleErr("PresentValue", err)
		}

		problem := &gen.StatusLog_Problem{
			Name:        fmt.Sprintf("%s:%s", o.Name, o.Test),
			Level:       gen.StatusLog_NOMINAL,
			Description: fmt.Sprintf("object %q nominal", o.Test),
		}
		level := gen.StatusLog_REDUCED_FUNCTION
		if o.Level != 0 {
			level = o.Level
		}

		switch {
		case notFoundCount == len(reqs):
			problem.Level = gen.StatusLog_NOTICE
			problem.Description = fmt.Sprintf("configured object %s not found", o.Test)
		case timeoutCount == len(reqs):
			problem.Level = level
			problem.Description = fmt.Sprintf("reading %s points %s timed out", o.Test, strings.Join(reqs, ", "))
		case len(errs) > 0:
			problem.Level = level
			problem.Description = fmt.Sprintf("error reading %s points %s", o.Test, strings.Join(reqs, ", "))
			m.Logger.Warn("device status: error reading properties",
				zap.String("test", o.Test),
				zap.Errors("errs", errs),
			)
		case hasValue && !nominalValue(valueVal, o.NominalValue):
			problem.Level = level
			problem.Description = fmt.Sprintf("read value %v: want %v, got %v", o.Test, o.NominalValue, valueVal)
		case hasStatus && (reliabilityVal != reliability.NoFaultDetected || eventStateVal != eventstate.Normal):
			var desc strings.Builder
			fmt.Fprintf(&desc, "object %v ", o.Test)
			var comma bool
			writeComma := func() {
				if comma {
					desc.WriteString(", ")
				}
				comma = true
			}
			if reliabilityVal != reliability.NoFaultDetected {
				writeComma()
				fmt.Fprintf(&desc, "reliability %v", reliabilityVal)
			}
			if eventStateVal != eventstate.Normal {
				writeComma()
				fmt.Fprintf(&desc, "event state %v", eventStateVal)
			}
			if outOfServiceVal {
				writeComma()
				fmt.Fprintf(&desc, "out of service")
			}
			if hasLimits {
				writeComma()
				fmt.Fprintf(&desc, "limits [%v,%v]", lowLimitVal, highLimitVal)
			}
			problem.Level = level
			problem.Description = desc.String()
		case !hasLimits:
			// m.Logger.Debug("device status",
			// 	zap.String("test", o.Test),
			// 	zap.Stringer("eventState", eventStateVal), zap.Stringer("reliability", reliabilityVal), zap.Bool("outOfService", outOfServiceVal),
			// )
		default:
			// m.Logger.Debug("device status",
			// 	zap.String("test", o.Test),
			// 	zap.Stringer("eventState", eventStateVal), zap.Stringer("reliability", reliabilityVal), zap.Bool("outOfService", outOfServiceVal),
			// 	zap.Bool("hasLimits", hasLimits), zap.Float32("lowLimit", lowLimitVal), zap.Float32("highLimit", highLimitVal),
			// )
		}

		m.status.UpdateProblem(o.Name, problem)
	}

	return nil
}

func nominalValue(v, nom any) bool {
	switch nv := nom.(type) {
	case float64: // json decodes numbers as float64 when the type is not specified
		switch vv := v.(type) {
		case int:
			return float64(vv) == nv
		case int8:
			return float64(vv) == nv
		case int16:
			return float64(vv) == nv
		case int32:
			return float64(vv) == nv
		case int64:
			return float64(vv) == nv
		case uint:
			return float64(vv) == nv
		case uint8:
			return float64(vv) == nv
		case uint16:
			return float64(vv) == nv
		case uint32:
			return float64(vv) == nv
		case uint64:
			return float64(vv) == nv
		case float32:
			return float64(vv) == nv
		case float64:
			return vv == nv
		case bool:
			if vv {
				return nv == 1
			} else {
				return nv == 0
			}
		default:
			return false
		}
	case bool:
		switch vv := v.(type) {
		case int:
			return (vv != 0) == nv
		case int8:
			return (vv != 0) == nv
		case int16:
			return (vv != 0) == nv
		case int32:
			return (vv != 0) == nv
		case int64:
			return (vv != 0) == nv
		case uint:
			return (vv != 0) == nv
		case uint8:
			return (vv != 0) == nv
		case uint16:
			return (vv != 0) == nv
		case uint32:
			return (vv != 0) == nv
		case uint64:
			return (vv != 0) == nv
		case float32:
			return (vv != 0) == nv
		case float64:
			return (vv != 0) == nv
		case bool:
			return vv == nv
		default:
			return false
		}
	default:
		return v == nom
	}
}

func asEnumType[T constraints.Unsigned](v any) (T, error) {
	value, err := comm.EnumValue(v)
	if err != nil {
		var zero T
		return zero, err
	}
	return T(value), nil
}
