package session

import "fmt"

type DeviceType int

const (
	DeviceTypeWeb DeviceType = iota
	DeviceTypeMobile
	DeviceTypeDesktop
	DeviceTypeTablet
)

var (
	DeviceTypeName = map[DeviceType]string{
		DeviceTypeWeb:     "Web",
		DeviceTypeMobile:  "Mobile",
		DeviceTypeDesktop: "Desktop",
		DeviceTypeTablet:  "Tablet",
	}
	DeviceTypeValue = map[string]DeviceType{
		"Web":     DeviceTypeWeb,
		"Mobile":  DeviceTypeMobile,
		"Desktop": DeviceTypeDesktop,
		"Tablet":  DeviceTypeTablet,
	}
)

func (t *DeviceType) String() string {
	return DeviceTypeName[*t]
}

func (t *DeviceType) IsValid() bool {
	switch *t {
	case DeviceTypeWeb, DeviceTypeMobile, DeviceTypeDesktop, DeviceTypeTablet:
		return true
	default:
		return false
	}
}

func (t *DeviceType) Scan(value any) error {
	i, ok := value.(int64)
	if !ok {
		return fmt.Errorf("cannot scan device type: %v", value)
	}

	*t = DeviceType(i)

	if !t.IsValid() {
		return fmt.Errorf("invalid device type: %d", i)
	}

	return nil
}
