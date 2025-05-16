package usecase

import (
	"context"
	"errors"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/dto"
	apperrors "github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/errors"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/port"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/infrastructure/data/spec"
)

type GetSessionsUseCase interface {
	UseCase[*dto.GetSessionsInputDTO, *dto.GetSessionsOutputDTO]
}

type GetSessionsUseCaseImpl struct {
	sessionRepo port.SessionsRepo
}

var _ GetSessionsUseCase = new(GetSessionsUseCaseImpl)

func NewGetSessionsUseCase(sessionRepo port.SessionsRepo) *GetSessionsUseCaseImpl {
	return &GetSessionsUseCaseImpl{
		sessionRepo: sessionRepo,
	}
}

func (u *GetSessionsUseCaseImpl) Execute(ctx context.Context, input *dto.GetSessionsInputDTO) (*dto.GetSessionsOutputDTO, error) {
	b := spec.NewSessionFilterSpec()

	s := b.
		WithUserId(input.UserId).
		WithIpAddress(input.IpAddr).
		WithDeviceInfo(input.DeviceInfo).
		WithIsActive(input.IsActive).
		WithExpiredBefore(input.ExpiredBefore).
		WithExpiredAfter(input.ExpiredAfter)

	sessions, err := u.sessionRepo.GetSessions(ctx, s)

	if errors.Is(err, context.DeadlineExceeded) {
		return nil, apperrors.NewDeadlineExceededError("getting sessions too long", err)
	}

	if errors.Is(err, context.Canceled) {
		return nil, apperrors.NewCanceledError("getting sessions closed", err)
	}

	if err != nil {
		return nil, apperrors.NewInternalError("something went wrong", err)
	}

	return &dto.GetSessionsOutputDTO{
		Sessions: sessions,
		Message:  "success",
	}, nil
}
