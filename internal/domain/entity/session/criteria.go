package session

import "time"

type Criteria struct {
	UserID           *int64
	OnlyActive       *bool
	DeviceType       *DeviceType
	DeviceName       *string
	AppType          *AppType
	AppVersion       *string
	OS               *string
	OSVersion        *string
	DeviceModel      *string
	IPAddr           *string
	LastActiveBefore *time.Time
	LastActiveAfter  *time.Time
	UpdatedBefore    *time.Time
	UpdatedAfter     *time.Time
	CreatedBefore    *time.Time
	CreatedAfter     *time.Time
}
