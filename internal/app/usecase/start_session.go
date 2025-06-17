package usecase

import (
	"context"
	"errors"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/dto"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/interfaces"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/pkg/apperrors"
)

type StartSession UseCase[*dto.StartSessionInputDTO, *dto.StartSessionOutputDTO]

type startSession struct {
	sessionService interfaces.SessionService
	sessionRepo    interfaces.SessionRepo
}

func NewStartSession(sessionService interfaces.SessionService, sessionRepo interfaces.SessionRepo) StartSession {
	return &startSession{
		sessionService: sessionService,
		sessionRepo:    sessionRepo,
	}
}

func (uc *startSession) Execute(ctx context.Context, input *dto.StartSessionInputDTO) (*dto.StartSessionOutputDTO, error) {
	if input == nil || input.DeviceInfo == nil {
		return nil, apperrors.NewInvalidArgumentError("Invalid input: missing session data", nil)
	}

	s := uc.buildSession(input)

	if err := s.Initialize(); err != nil {
		return nil, apperrors.NewInternalError("Unexpected server error. Please try again.").WithCause(err)
	}

	uc.sessionService.EnrichLocation(s)

	sessionID, err := uc.sessionRepo.Create(ctx, s)

	if err != nil {
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			return nil, apperrors.NewDeadlineExceededError("Session creation timed out. Please try again.").WithCause(err)
		case errors.Is(err, context.Canceled):
			return nil, apperrors.NewCanceledError("Session creation was canceled.").WithCause(err)
		default:
			return nil, apperrors.NewInternalError("Failed to create session.").WithCause(err)
		}
	}

	return &dto.StartSessionOutputDTO{
		SessionID: *sessionID,
		Message:   "Session created",
	}, nil
}

func (uc *startSession) buildSession(input *dto.StartSessionInputDTO) *session.Session {
	b := session.NewBuilder()

	return b.
		WithUserID(input.UserID).
		WithDeviceInfo(input.DeviceInfo).
		Build()
}
