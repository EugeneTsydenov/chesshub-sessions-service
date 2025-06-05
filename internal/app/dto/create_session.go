package dto

import (
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
	"github.com/google/uuid"
)

type (
	CreateSessionInputDTO struct {
		UserID     int64
		DeviceInfo *session.DeviceInfo
	}

	CreateSessionOutputDTO struct {
		SessionID uuid.UUID
		Message   string
	}
)
