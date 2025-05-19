package dto

import (
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity"
	"time"
)

type (
	GetSessionsInputDTO struct {
		UserID        *int64
		IPAddr        *string
		DeviceInfo    *string
		IsActive      *bool
		ExpiredAfter  time.Time
		ExpiredBefore time.Time
	}

	GetSessionsOutputDTO struct {
		Sessions []*entity.Session
		Message  string
	}
)
