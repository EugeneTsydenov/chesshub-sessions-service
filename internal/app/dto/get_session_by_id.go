package dto

import "github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity"

type (
	GetSessionByIDInputDTO struct {
		ID string
	}

	GetSessionByIDOutputDTO struct {
		Session *entity.Session
		Message string
	}
)
