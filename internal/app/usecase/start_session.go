package usecase

import (
	"context"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/dto"
	apperrors "github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/errors"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/interfaces"
)

type (
	StartSession UseCase[*dto.StartSessionInputDTO, *dto.StartSessionOutputDTO]

	startSession struct {
		sessionService interfaces.SessionService
		sessionRepo    interfaces.SessionRepo
	}
)

var _ StartSession = new(startSession)

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
		return nil, apperrors.NewInternalError("Unexpected server error.").WithCause(err)
	}

	uc.sessionService.EnrichLocation(s)

	sessionID, err := uc.sessionRepo.Create(ctx, s)
	if err != nil {
		return nil, apperrors.FromDomainError(err)
	}

	return &dto.StartSessionOutputDTO{
		SessionID: sessionID,
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
