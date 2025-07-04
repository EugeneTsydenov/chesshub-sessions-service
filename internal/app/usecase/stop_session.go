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
	StopSession UseCase[*dto.StopSessionInputDTO, *dto.StopSessionOutputDTO]

	stopSession struct {
		sessionService interfaces.SessionService
		sessionRepo    interfaces.SessionRepo
	}
)

var _ StopSession = new(stopSession)

func NewStopSession(sessionService interfaces.SessionService, repo interfaces.SessionRepo) StopSession {
	return &stopSession{
		sessionService: sessionService,
		sessionRepo:    repo,
	}
}

func (uc *stopSession) Execute(ctx context.Context, input *dto.StopSessionInputDTO) (*dto.StopSessionOutputDTO, error) {
	sessionID := input.SessionID

	sessionUUID, err := uuid.Parse(sessionID)
	if err != nil {
		return nil, apperrors.NewInvalidArgumentError(fmt.Sprintf("Invalid session id: %s", sessionID), nil).WithCause(err)
	}

	s, err := uc.sessionRepo.GetByID(ctx, sessionUUID)
	if err != nil {
		return nil, apperrors.FromDomainError(err)
	}

	err = uc.sessionService.DeactivateSession(ctx, s)
	if err != nil {
		return nil, apperrors.FromDomainError(err)
	}

	return &dto.StopSessionOutputDTO{
		Message: fmt.Sprintf("Session with ID %s has been deactivated", sessionID),
	}, nil
}
