package usecase

import (
	"context"
	"fmt"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/dto"
	apperrors "github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/errors"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/interfaces"
	"github.com/google/uuid"
)

type (
	GetSession UseCase[*dto.GetSessionInputDTO, *dto.GetSessionOutputDTO]

	getSession struct {
		cachedSessionRepo interfaces.SessionRepo
	}
)

var _ GetSession = new(getSession)

func NewGetSession(cachedRepo interfaces.SessionRepo) GetSession {
	return &getSession{
		cachedSessionRepo: cachedRepo,
	}
}

func (uc getSession) Execute(ctx context.Context, input *dto.GetSessionInputDTO) (*dto.GetSessionOutputDTO, error) {
	sessionID := input.SessionID

	sessionUUID, err := uuid.Parse(sessionID)
	if err != nil {
		return nil, apperrors.NewInvalidArgumentError(fmt.Sprintf("Invalid session id: %s", sessionID), nil).WithCause(err)
	}

	s, err := uc.cachedSessionRepo.GetByID(ctx, sessionUUID)
	if err != nil {
		return nil, apperrors.FromDomainError(err)
	}

	return &dto.GetSessionOutputDTO{
		Session: s,
		Message: "Session was got",
	}, nil
}
