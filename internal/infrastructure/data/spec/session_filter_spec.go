package spec

import (
	"fmt"
	"strings"
	"time"
)

type SessionFilerSpec interface {
	Spec
}

type SessionFilterSpecImpl struct {
	UserId        *int64
	IpAddress     *string
	DeviceInfo    *string
	IsActive      *bool
	ExpiredAfter  time.Time
	ExpiredBefore time.Time
}

var _ SessionFilerSpec = new(SessionFilterSpecImpl)

func NewSessionFilterSpec() *SessionFilterSpecImpl {
	return &SessionFilterSpecImpl{}
}

func (f *SessionFilterSpecImpl) WithUserId(userId *int64) *SessionFilterSpecImpl {
	f.UserId = userId
	return f
}

func (f *SessionFilterSpecImpl) WithIpAddress(ipAddress *string) *SessionFilterSpecImpl {
	f.IpAddress = ipAddress
	return f
}

func (f *SessionFilterSpecImpl) WithDeviceInfo(deviceInfo *string) *SessionFilterSpecImpl {
	f.DeviceInfo = deviceInfo
	return f
}

func (f *SessionFilterSpecImpl) WithIsActive(isActive *bool) *SessionFilterSpecImpl {
	f.IsActive = isActive
	return f
}

func (f *SessionFilterSpecImpl) WithExpiredAfter(expiredAfter time.Time) *SessionFilterSpecImpl {
	f.ExpiredAfter = expiredAfter
	return f
}

func (f *SessionFilterSpecImpl) WithExpiredBefore(expiredBefore time.Time) *SessionFilterSpecImpl {
	f.ExpiredBefore = expiredBefore
	return f
}

func (f *SessionFilterSpecImpl) BuildQuery() (string, []any) {
	baseQuery := "SELECT * FROM sessions"
	var args []any
	var conditions []string

	paramCount := 1

	if f.UserId != nil {
		conditions = append(conditions, fmt.Sprintf("user_id = $%d", paramCount))
		args = append(args, *f.UserId)
		paramCount++
	}

	if f.IpAddress != nil {
		conditions = append(conditions, fmt.Sprintf("ip_address = $%d", paramCount))
		args = append(args, *f.IpAddress)
		paramCount++
	}

	if f.DeviceInfo != nil {
		conditions = append(conditions, fmt.Sprintf("device_info = $%d", paramCount))
		args = append(args, *f.DeviceInfo)
		paramCount++
	}

	if f.IsActive != nil {
		conditions = append(conditions, fmt.Sprintf("is_active = $%d", paramCount))
		args = append(args, *f.IsActive)
		paramCount++
	}

	if !f.ExpiredAfter.IsZero() {
		conditions = append(conditions, fmt.Sprintf("expired_at > $%d", paramCount))
		args = append(args, f.ExpiredAfter)
		paramCount++
	}

	if !f.ExpiredBefore.IsZero() {
		conditions = append(conditions, fmt.Sprintf("expired_at < $%d", paramCount))
		args = append(args, f.ExpiredBefore)
		paramCount++
	}

	if len(conditions) > 0 {
		baseQuery += " WHERE " + strings.Join(conditions, " AND ")
	}

	return baseQuery, args
}
