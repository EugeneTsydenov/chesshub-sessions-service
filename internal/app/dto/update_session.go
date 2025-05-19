package dto

import (
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity"
	"time"
)

type (
	UpdateSessionInputDTO struct {
		SessionID  string
		IpAddr     *string
		DeviceInfo *string
		IsActive   *bool
		ExpiredAt  time.Time
	}

	UpdateSessionOutputDTO struct {
		Session *entity.Session
		Message string
	}
)
