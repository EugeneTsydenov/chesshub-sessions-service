package dto

import (
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
)

type (
	GetSessionInputDTO struct {
		SessionID string
	}

	GetSessionOutputDTO struct {
		Session *session.Session
		Message string
	}
)
