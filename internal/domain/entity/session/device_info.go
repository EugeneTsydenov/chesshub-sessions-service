package session

type DeviceInfo struct {
	deviceType  DeviceType
	deviceName  *string
	appType     AppType
	appVersion  string
	os          string
	osVersion   *string
	deviceModel *string
	ipAddr      string
}

func NewDeviceInfo(
	deviceType DeviceType,
	deviceName *string,
	appType AppType,
	appVersion,
	os string,
	osVersion,
	deviceModel *string,
	ipAddr string,
) *DeviceInfo {
	return &DeviceInfo{
		deviceType:  deviceType,
		deviceName:  deviceName,
		appType:     appType,
		appVersion:  appVersion,
		os:          os,
		osVersion:   osVersion,
		deviceModel: deviceModel,
		ipAddr:      ipAddr,
	}
}

func (d *DeviceInfo) DeviceType() DeviceType {
	return d.deviceType
}

func (d *DeviceInfo) DeviceName() *string {
	return d.deviceName
}

func (d *DeviceInfo) AppType() AppType {
	return d.appType
}

func (d *DeviceInfo) AppVersion() string {
	return d.appVersion
}

func (d *DeviceInfo) OS() string {
	return d.os
}

func (d *DeviceInfo) OSVersion() *string {
	return d.osVersion
}

func (d *DeviceInfo) DeviceModel() *string {
	return d.deviceModel
}

func (d *DeviceInfo) IPAddr() string {
	return d.ipAddr
}
