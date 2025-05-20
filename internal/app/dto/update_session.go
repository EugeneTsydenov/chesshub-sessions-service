package dto

import (
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity"
)

type (
	UpdateSessionInputDTO struct {
		SessionID string
		FieldMap  map[string]string
	}

	UpdateSessionOutputDTO struct {
		Session *entity.Session
		Message string
	}
)
