package adapt

import (
	"fmt"

	"github.com/vanti-dev/gobacnet"
	bactypes "github.com/vanti-dev/gobacnet/types"
	"github.com/vanti-dev/gobacnet/types/objecttype"
	"github.com/vanti-dev/sc-bos/pkg/gentrait/statuspb"

	"github.com/vanti-dev/sc-bos/pkg/driver/bacnet/config"
	"github.com/vanti-dev/sc-bos/pkg/driver/bacnet/rpc"
	"github.com/vanti-dev/sc-bos/pkg/node"
)

// Object adapts a bacnet object into one or more smart core named traits.
func Object(prefix string, client *gobacnet.Client, device bactypes.Device, object config.Object, statuses *statuspb.Map) (node.SelfAnnouncer, error) {
	switch object.ID.Type {
	case objecttype.BinaryValue, objecttype.BinaryOutput, objecttype.BinaryInput:
		return BinaryObject(prefix, client, device, object, statuses)
	}

	if object.Trait == "" {
		return nil, ErrNoDefault
	}
	return nil, ErrNoAdaptation
}

// DeviceName returns the smart core name we should use for the configured object.
func DeviceName(o config.Device) string {
	if o.Name != "" {
		return o.Name
	}
	return fmt.Sprintf("%d", o.ID)
}

// ObjectName returns the smart core name we should use for the configured object.
func ObjectName(o config.Object) string {
	if o.Name != "" {
		return o.Name
	}
	return o.ID.String()
}

func ObjectIDFromProto(identifier *rpc.ObjectIdentifier) bactypes.ObjectID {
	return bactypes.ObjectID{
		Type:     objecttype.ObjectType(identifier.Type),
		Instance: bactypes.ObjectInstance(identifier.Instance),
	}
}

func ObjectIDToProto(id bactypes.ObjectID) *rpc.ObjectIdentifier {
	return &rpc.ObjectIdentifier{
		Type:     uint32(id.Type),
		Instance: uint32(id.Instance),
	}
}
