package session

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
