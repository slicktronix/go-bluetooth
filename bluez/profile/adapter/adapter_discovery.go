package adapter

import (
	"sync"

	"github.com/godbus/dbus/v5"
	log "github.com/sirupsen/logrus"
	"github.com/slicktronix/go-bluetooth/bluez"
	"github.com/slicktronix/go-bluetooth/bluez/profile/device"
)

const (
	// DeviceRemoved a device has been removed from local cache
	DeviceRemoved DeviceActions = iota
	// DeviceAdded new device found, eg. via discovery
	DeviceAdded
)

type DeviceActions uint8

// DeviceDiscovered event emitted when a device is added or removed from Object Manager
type DeviceDiscovered struct {
	Path dbus.ObjectPath
	Type DeviceActions
}

// OnDeviceDiscovered monitor for new devices and send updates via channel. Use cancel to close the monitoring process
func (a *Adapter1) OnDeviceDiscovered() (chan *DeviceDiscovered, func(), error) {

	signal, omSignalCancel, err := a.GetObjectManagerSignal()
	if err != nil {
		return nil, nil, err
	}

	var (
		ch    = make(chan *DeviceDiscovered)
		mutex sync.Mutex
	)

	go func() {
		// Recover from panic on write to closed channel which happens
		// very often when there's too many BLE advertisements to process
		// in timly manner by bluez+dbus and advertising reports come in
		// after scanning was stopped
		defer func() {
			if err := recover(); err != nil {
				log.Warnf("Recovering from panic: %s", err)
			}
		}()

		for v := range signal {

			if v == nil {
				return
			}

			var op DeviceActions
			if v.Name == bluez.InterfacesAdded {
				op = DeviceAdded
			} else {
				if v.Name == bluez.InterfacesRemoved {
					op = DeviceRemoved
				} else {
					continue
				}
			}

			path := v.Body[0].(dbus.ObjectPath)

			if op == DeviceRemoved {
				ifaces := v.Body[1].([]string)
				for _, iface := range ifaces {
					if iface == device.Device1Interface {
						log.Tracef("Removed device %s", path)
						mutex.Lock()
						ch <- &DeviceDiscovered{path, op}
						mutex.Unlock()
					}
				}
				continue
			}

			ifaces := v.Body[1].(map[string]map[string]dbus.Variant)
			if p, ok := ifaces[device.Device1Interface]; ok {
				if p == nil {
					continue
				}
				log.Tracef("Added device %s", path)

				mutex.Lock()
				if ch == nil {
					return
				}

				ch <- &DeviceDiscovered{path, op}
				mutex.Unlock()
			}

		}
	}()

	cancel := func() {
		omSignalCancel()
		mutex.Lock()
		if ch != nil {
			close(ch)
		}
		ch = nil
		mutex.Unlock()
		log.Trace("OnDeviceDiscovered: cancel() called")
	}

	return ch, cancel, nil
}
