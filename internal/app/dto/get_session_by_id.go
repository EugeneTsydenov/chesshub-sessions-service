package dto

import "github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity"

type (
	GetSessionByIdInputDTO struct {
		ID string
	}

	GetSessionByIdOutputDTO struct {
		Session *entity.Session
		Message string
	}
)
