package dto

import (
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/sessionfilter"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
)

type (
	ListSessionsInputDTO struct {
		Filter *sessionfilter.SessionFilter
	}

	ListSessionsOutputDTO struct {
		Sessions []*session.Session
		Count    int
		Message  string
	}
)
