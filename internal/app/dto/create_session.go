package dto

import (
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity"
	"time"
)

type (
	CreateSessionInputDTO struct {
		UserId     int64
		IpAddr     string
		DeviceInfo string
		ExpiredAt  time.Time
	}

	CreateSessionOutputDTO struct {
		Session *entity.Session
		Message string
	}
)
