package usecase

import (
	"context"
	"errors"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/dto"
	apperrors "github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/errors"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/port"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity"
)

type CreateSessionUseCase interface {
	UseCase[*dto.CreateSessionInputDTO, *dto.CreateSessionOutputDTO]
}

type CreateSessionUseCaseImpl struct {
	sessionsRepo port.SessionsRepo
}

var _ CreateSessionUseCase = new(CreateSessionUseCaseImpl)

func NewCreateSessionUseCase(sessionsRepo port.SessionsRepo) *CreateSessionUseCaseImpl {
	return &CreateSessionUseCaseImpl{
		sessionsRepo: sessionsRepo,
	}
}

func (u *CreateSessionUseCaseImpl) Execute(ctx context.Context, input *dto.CreateSessionInputDTO) (*dto.CreateSessionOutputDTO, error) {
	b := entity.NewSessionBuilder()

	s := b.
		WithUserID(input.UserID).
		WithIPAddr(input.IPAddr).
		WithDeviceInfo(input.DeviceInfo).
		WithExpiredAt(input.ExpiredAt).
		Build()

	createdSession, err := u.sessionsRepo.Create(ctx, s)

	if errors.Is(err, context.DeadlineExceeded) {
		return nil, apperrors.NewDeadlineExceededError("creation session too long", err)
	}

	if errors.Is(err, context.Canceled) {
		return nil, apperrors.NewCanceledError("creation session closed", err)
	}

	if err != nil {
		return nil, apperrors.NewInternalError("Unexpected internal error", err)
	}

	return &dto.CreateSessionOutputDTO{
		Session: createdSession,
		Message: "Session created",
	}, nil
}
