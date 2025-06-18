package usecase

import (
	"context"
	"fmt"

	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/dto"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/interfaces"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/pkg/apperrors"
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

func NewStopSession(sessionService interfaces.SessionService, sessionRepo interfaces.SessionRepo) StopSession {
	return &stopSession{
		sessionService: sessionService,
		sessionRepo:    sessionRepo,
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
		return nil, apperrors.NewInternalError("Failed to stop session.").WithCause(err)
	}

	if s.IsEmpty() {
		return nil, apperrors.NewNotFoundError(fmt.Sprintf("Session with ID %s not found", sessionID))
	}

	err = uc.sessionService.DeactivateSession(ctx, s)
	if err != nil {
		return nil, apperrors.NewInternalError("Failed to stop session.").WithCause(err)
	}

	return &dto.StopSessionOutputDTO{
		Message: fmt.Sprintf("Session with ID %s has been deactivated", sessionID),
	}, nil
}
