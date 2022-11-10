package bacnet

import (
	"fmt"
	"github.com/vanti-dev/bsp-ew/internal/driver/bacnet/config"
	bactypes "github.com/vanti-dev/gobacnet/types"
	"net"
)

func (d *Driver) findDevice(device config.Device) (bactypes.Device, error) {
	fail := func(err error) (bactypes.Device, error) {
		return bactypes.Device{}, err
	}

	if device.Comm == nil {
		id := device.ID
		is, err := d.client.WhoIs(int(id), int(id))
		if err != nil {
			return fail(err)
		}
		if len(is) == 0 {
			return fail(fmt.Errorf("no devices found (via WhoIs) with id %d", id))
		}
		return is[0], nil
	}

	udpAddr := net.UDPAddrFromAddrPort(*device.Comm.IP)
	addr := bactypes.UDPToAddress(udpAddr)
	bacDevices, err := d.client.RemoteDevices(addr, device.ID)
	if err != nil {
		return fail(err)
	}
	return bacDevices[0], nil
}

func (d *Driver) fetchObjects(cfg config.Root, device config.Device, bacDevice bactypes.Device) (map[bactypes.ObjectID]configObject, error) {
	objects := make(map[bactypes.ObjectID]configObject, len(device.Objects))
	for _, object := range device.Objects {
		objects[bactypes.ObjectID(object.ID)] = configObject{
			co: object,
			bo: &bactypes.Object{
				ID: bactypes.ObjectID(object.ID),
			},
		}
	}

	discoverObjects := cfg.DiscoverObjects
	if device.DiscoverObjects != nil {
		discoverObjects = *device.DiscoverObjects
	}

	if discoverObjects {
		hasObjects, err := d.client.Objects(bacDevice)
		if err != nil {
			return objects, fmt.Errorf("read objects %w", err)
		}

		for _, objectsOfType := range hasObjects.Objects {
			for _, object := range objectsOfType {
				object := object
				if known, found := objects[object.ID]; found {
					// copy any additional data into the object config
					if known.co.Title == "" {
						known.co.Title = firstNonEmpty(object.Description, object.Name)
					}
					known.bo = &object
					objects[object.ID] = known
					continue
				}
				objects[object.ID] = configObject{
					co: config.Object{
						ID:    config.ObjectID(object.ID),
						Title: firstNonEmpty(object.Description, object.Name),
					},
					bo: &object,
				}
			}
		}
	}

	return objects, nil
}

type configObject struct {
	co config.Object
	bo *bactypes.Object
}

func firstNonEmpty(strs ...string) string {
	for _, str := range strs {
		if str != "" {
			return str
		}
	}
	return ""
}
