package dto

import (
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity"
)

type (
	GetSessionsInputDTO struct {
		FilterMap map[string]string
	}

	GetSessionsOutputDTO struct {
		Sessions []*entity.Session
		Message  string
	}
)
