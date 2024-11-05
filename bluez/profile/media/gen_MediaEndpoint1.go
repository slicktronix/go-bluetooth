// Code generated by go-bluetooth generator DO NOT EDIT.

package media

import (
	"sync"

	"github.com/godbus/dbus/v5"
	"github.com/slicktronix/go-bluetooth/bluez"
	"github.com/slicktronix/go-bluetooth/props"
	"github.com/slicktronix/go-bluetooth/util"
)

var MediaEndpoint1Interface = "org.bluez.MediaEndpoint1"

// NewMediaEndpoint1Server create a new instance of MediaEndpoint1
//
// Args:
// - servicePath: unique name
func NewMediaEndpoint1Server(servicePath string, objectPath dbus.ObjectPath) (*MediaEndpoint1, error) {
	a := new(MediaEndpoint1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  servicePath,
			Iface: MediaEndpoint1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	a.Properties = new(MediaEndpoint1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	return a, nil
}

// NewMediaEndpoint1Client create a new instance of MediaEndpoint1
//
// Args:

func NewMediaEndpoint1Client(objectPath dbus.ObjectPath) (*MediaEndpoint1, error) {
	a := new(MediaEndpoint1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez",
			Iface: MediaEndpoint1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	a.Properties = new(MediaEndpoint1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	return a, nil
}

/*
MediaEndpoint1 MediaEndpoint1 hierarchy
*/
type MediaEndpoint1 struct {
	client                 *bluez.Client
	propertiesSignal       chan *dbus.Signal
	objectManagerSignal    chan *dbus.Signal
	objectManager          *bluez.ObjectManager
	Properties             *MediaEndpoint1Properties
	watchPropertiesChannel chan *dbus.Signal
}

// MediaEndpoint1Properties contains the exposed properties of an interface
type MediaEndpoint1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`

	/*
		Capabilities Capabilities blob, it is used as it is so the size and
				byte order must match.
	*/
	Capabilities []byte

	/*
		Codec Assigned number of codec that the endpoint implements.
				The values should match the profile specification which
				is indicated by the UUID.
	*/
	Codec byte

	/*
		Device Device object which the endpoint is belongs to.
	*/
	Device dbus.ObjectPath

	/*
		UUID UUID of the profile which the endpoint is for.
	*/
	UUID string
}

// Lock access to properties
func (p *MediaEndpoint1Properties) Lock() {
	p.lock.Lock()
}

// Unlock access to properties
func (p *MediaEndpoint1Properties) Unlock() {
	p.lock.Unlock()
}

// SetCapabilities set Capabilities value
func (a *MediaEndpoint1) SetCapabilities(v []byte) error {
	return a.SetProperty("Capabilities", v)
}

// GetCapabilities get Capabilities value
func (a *MediaEndpoint1) GetCapabilities() ([]byte, error) {
	v, err := a.GetProperty("Capabilities")
	if err != nil {
		return []byte{}, err
	}
	return v.Value().([]byte), nil
}

// SetCodec set Codec value
func (a *MediaEndpoint1) SetCodec(v byte) error {
	return a.SetProperty("Codec", v)
}

// GetCodec get Codec value
func (a *MediaEndpoint1) GetCodec() (byte, error) {
	v, err := a.GetProperty("Codec")
	if err != nil {
		return byte(0), err
	}
	return v.Value().(byte), nil
}

// SetDevice set Device value
func (a *MediaEndpoint1) SetDevice(v dbus.ObjectPath) error {
	return a.SetProperty("Device", v)
}

// GetDevice get Device value
func (a *MediaEndpoint1) GetDevice() (dbus.ObjectPath, error) {
	v, err := a.GetProperty("Device")
	if err != nil {
		return dbus.ObjectPath(""), err
	}
	return v.Value().(dbus.ObjectPath), nil
}

// SetUUID set UUID value
func (a *MediaEndpoint1) SetUUID(v string) error {
	return a.SetProperty("UUID", v)
}

// GetUUID get UUID value
func (a *MediaEndpoint1) GetUUID() (string, error) {
	v, err := a.GetProperty("UUID")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}

// Close the connection
func (a *MediaEndpoint1) Close() {
	a.unregisterPropertiesSignal()
	a.client.Disconnect()
}

// Path return MediaEndpoint1 object path
func (a *MediaEndpoint1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Client return MediaEndpoint1 dbus client
func (a *MediaEndpoint1) Client() *bluez.Client {
	return a.client
}

// Interface return MediaEndpoint1 interface
func (a *MediaEndpoint1) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *MediaEndpoint1) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

	if a.objectManagerSignal == nil {
		if a.objectManager == nil {
			om, err := bluez.GetObjectManager()
			if err != nil {
				return nil, nil, err
			}
			a.objectManager = om
		}

		s, err := a.objectManager.Register()
		if err != nil {
			return nil, nil, err
		}
		a.objectManagerSignal = s
	}

	cancel := func() {
		if a.objectManagerSignal == nil {
			return
		}
		a.objectManagerSignal <- nil
		a.objectManager.Unregister(a.objectManagerSignal)
		a.objectManagerSignal = nil
	}

	return a.objectManagerSignal, cancel, nil
}

// ToMap convert a MediaEndpoint1Properties to map
func (a *MediaEndpoint1Properties) ToMap() (map[string]interface{}, error) {
	return props.ToMap(a), nil
}

// FromMap convert a map to an MediaEndpoint1Properties
func (a *MediaEndpoint1Properties) FromMap(props map[string]interface{}) (*MediaEndpoint1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an MediaEndpoint1Properties
func (a *MediaEndpoint1Properties) FromDBusMap(props map[string]dbus.Variant) (*MediaEndpoint1Properties, error) {
	s := new(MediaEndpoint1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// ToProps return the properties interface
func (a *MediaEndpoint1) ToProps() bluez.Properties {
	return a.Properties
}

// GetWatchPropertiesChannel return the dbus channel to receive properties interface
func (a *MediaEndpoint1) GetWatchPropertiesChannel() chan *dbus.Signal {
	return a.watchPropertiesChannel
}

// SetWatchPropertiesChannel set the dbus channel to receive properties interface
func (a *MediaEndpoint1) SetWatchPropertiesChannel(c chan *dbus.Signal) {
	a.watchPropertiesChannel = c
}

// GetProperties load all available properties
func (a *MediaEndpoint1) GetProperties() (*MediaEndpoint1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *MediaEndpoint1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *MediaEndpoint1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *MediaEndpoint1) GetPropertiesSignal() (chan *dbus.Signal, error) {

	if a.propertiesSignal == nil {
		s, err := a.client.Register(a.client.Config.Path, bluez.PropertiesInterface)
		if err != nil {
			return nil, err
		}
		a.propertiesSignal = s
	}

	return a.propertiesSignal, nil
}

// Unregister for changes signalling
func (a *MediaEndpoint1) unregisterPropertiesSignal() {
	if a.propertiesSignal != nil {
		a.propertiesSignal <- nil
		a.propertiesSignal = nil
	}
}

// WatchProperties updates on property changes
func (a *MediaEndpoint1) WatchProperties() (chan *bluez.PropertyChanged, error) {
	return bluez.WatchProperties(a)
}

func (a *MediaEndpoint1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	return bluez.UnwatchProperties(a, ch)
}

/*
SetConfiguration 			Set configuration for the transport.

	For client role transport must be set with a server
	endpoint oject which will be configured and the
	properties must contain the following properties:
		array{byte} Capabilities
*/
func (a *MediaEndpoint1) SetConfiguration(transport dbus.ObjectPath, properties map[string]interface{}) error {
	return a.client.Call("SetConfiguration", 0, transport, properties).Store()
}

/*
SelectConfiguration 			Select preferable configuration from the supported

	capabilities.
	Returns a configuration which can be used to setup
	a transport.
	Note: There is no need to cache the selected
	configuration since on success the configuration is
	send back as parameter of SetConfiguration.
*/
func (a *MediaEndpoint1) SelectConfiguration(capabilities []byte) ([]byte, error) {
	val0 := []byte{}
	err := a.client.Call("SelectConfiguration", 0, capabilities).Store(&val0)
	return val0, err
}

/*
ClearConfiguration 			Clear transport configuration.
*/
func (a *MediaEndpoint1) ClearConfiguration(transport dbus.ObjectPath) error {
	return a.client.Call("ClearConfiguration", 0, transport).Store()
}

/*
Release 			This method gets called when the service daemon

	unregisters the endpoint. An endpoint can use it to do
	cleanup tasks. There is no need to unregister the
	endpoint, because when this method gets called it has
	already been unregistered.
*/
func (a *MediaEndpoint1) Release() error {
	return a.client.Call("Release", 0).Store()
}
