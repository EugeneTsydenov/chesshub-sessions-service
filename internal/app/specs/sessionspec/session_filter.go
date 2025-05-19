package sessionspec

import (
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/pkg/spec"
	"time"
)

type SessionFilterSpec struct {
	userID        *int64
	ipAddr        *string
	deviceInfo    *string
	isActive      *bool
	expiredAfter  time.Time
	expiredBefore time.Time
}

func NewSessionFilterSpec(userID *int64, ipAddr, deviceInfo *string, isActive *bool, expiredAfter, expiredBefore time.Time) *SessionFilterSpec {
	return &SessionFilterSpec{
		userID:        userID,
		ipAddr:        ipAddr,
		deviceInfo:    deviceInfo,
		isActive:      isActive,
		expiredAfter:  expiredAfter,
		expiredBefore: expiredBefore,
	}
}

func (s *SessionFilterSpec) ToSQL() (string, []any) {
	baseQuery := "SELECT * FROM sessions"
	var specs []spec.Spec
	paramIndex := 1

	if s.userID != nil {
		specs = append(specs, spec.NewFieldSpec("user_id", "=", paramIndex, *s.userID))
		paramIndex++
	}

	if s.ipAddr != nil {
		specs = append(specs, spec.NewFieldSpec("ip_address", "=", paramIndex, *s.ipAddr))
		paramIndex++
	}

	if s.deviceInfo != nil {
		specs = append(specs, spec.NewFieldSpec("device_info", "=", paramIndex, *s.deviceInfo))
		paramIndex++
	}

	if s.isActive != nil {
		specs = append(specs, spec.NewFieldSpec("is_active", "=", paramIndex, *s.isActive))
		paramIndex++
	}

	if !s.expiredAfter.IsZero() {
		specs = append(specs, spec.NewFieldSpec("expired_at", ">", paramIndex, s.expiredAfter))
		paramIndex++
	}

	if !s.expiredBefore.IsZero() {
		specs = append(specs, spec.NewFieldSpec("expired_at", "<", paramIndex, s.expiredBefore))
		paramIndex++
	}

	andSpec := spec.NewAnd(specs...)

	query, args := andSpec.ToSQL()

	return baseQuery + " WHERE " + query, args
}
