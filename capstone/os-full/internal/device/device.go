package device

import (
	"fmt"
	"sync"
)

type DeviceType int

const (
	TypeDisk DeviceType = iota
	TypeConsole
	TypeNetwork
)

type Device struct {
	Name      string
	Type      DeviceType
	Major     int
	Minor     int
	Mounted   bool
	MountPath string
	Mu        sync.RWMutex
}

type DeviceManager struct {
	devices map[string]*Device
	Mu      sync.RWMutex
}

func NewDeviceManager() *DeviceManager {
	return &DeviceManager{
		devices: make(map[string]*Device),
	}
}

func (dm *DeviceManager) RegisterDevice(name string, devType DeviceType, major, minor int) *Device {
	dm.Mu.Lock()
	defer dm.Mu.Unlock()

	device := &Device{
		Name:    name,
		Type:    devType,
		Major:   major,
		Minor:   minor,
		Mounted: false,
	}
	dm.devices[name] = device
	return device
}

func (dm *DeviceManager) GetDevice(name string) (*Device, error) {
	dm.Mu.RLock()
	defer dm.Mu.RUnlock()

	device, exists := dm.devices[name]
	if !exists {
		return nil, fmt.Errorf("device not found: %s", name)
	}
	return device, nil
}

func (dm *DeviceManager) MountDevice(name string, path string) error {
	dm.Mu.Lock()
	defer dm.Mu.Unlock()

	device, exists := dm.devices[name]
	if !exists {
		return fmt.Errorf("device not found: %s", name)
	}

	device.Mu.Lock()
	defer device.Mu.Unlock()

	if device.Mounted {
		return fmt.Errorf("device already mounted: %s", name)
	}

	device.Mounted = true
	device.MountPath = path
	return nil
}

func (dm *DeviceManager) UnmountDevice(name string) error {
	dm.Mu.Lock()
	defer dm.Mu.Unlock()

	device, exists := dm.devices[name]
	if !exists {
		return fmt.Errorf("device not found: %s", name)
	}

	device.Mu.Lock()
	defer device.Mu.Unlock()

	if !device.Mounted {
		return fmt.Errorf("device not mounted: %s", name)
	}

	device.Mounted = false
	device.MountPath = ""
	return nil
}

func (dm *DeviceManager) ListDevices() []*Device {
	dm.Mu.RLock()
	defer dm.Mu.RUnlock()

	devices := make([]*Device, 0, len(dm.devices))
	for _, device := range dm.devices {
		devices = append(devices, device)
	}
	return devices
}

func (d *Device) String() string {
	d.Mu.RLock()
	defer d.Mu.RUnlock()

	typeStr := "Unknown"
	switch d.Type {
	case TypeDisk:
		typeStr = "Disk"
	case TypeConsole:
		typeStr = "Console"
	case TypeNetwork:
		typeStr = "Network"
	}

	mounted := "No"
	if d.Mounted {
		mounted = d.MountPath
	}

	return fmt.Sprintf("%s (%s) - Major: %d, Minor: %d, Mounted: %s",
		d.Name, typeStr, d.Major, d.Minor, mounted)
}
