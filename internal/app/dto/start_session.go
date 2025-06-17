package dto

import (
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
	"github.com/google/uuid"
)

type (
	StartSessionInputDTO struct {
		UserID     int64
		DeviceInfo *session.DeviceInfo
	}

	StartSessionOutputDTO struct {
		SessionID uuid.UUID
		Message   string
	}
)
