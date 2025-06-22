package sessionfilter

import (
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
	"time"
)

type SessionFilter struct {
	UserID           *int64
	OnlyActive       *bool
	DeviceType       *session.DeviceType
	DeviceName       *string
	AppType          *session.AppType
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
