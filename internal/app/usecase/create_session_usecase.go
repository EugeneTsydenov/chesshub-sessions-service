package usecase

import (
	"context"
	"errors"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/apperrors"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/dto"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/infra/session"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/infra/session/data/repo"
	"log"
	"time"
)

type CreateSessionUseCase interface {
	Executor[*dto.CreateSessionInputDto, *dto.CreateSessionOutputDto]
}

type CreateSessionUseCaseImpl struct {
	sessionsRepo repo.SessionsRepo
}

var _ CreateSessionUseCase = new(CreateSessionUseCaseImpl)

func NewCreateSessionUseCase(sessionRepo repo.SessionsRepo) *CreateSessionUseCaseImpl {
	return &CreateSessionUseCaseImpl{
		sessionsRepo: sessionRepo,
	}
}

func (u *CreateSessionUseCaseImpl) Execute(ctx context.Context, input *dto.CreateSessionInputDto) (*dto.CreateSessionOutputDto, error) {
	b := session.NewBuilder()

	s := b.
		WithUserId(input.UserId).
		WithIpAddr(input.IpAddr).
		WithDeviceInfo(input.DeviceInfo).
		WithExpiredAt(input.ExpiredAt).
		Build()

	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	createdSession, err := u.sessionsRepo.Create(ctx, s)

	if errors.Is(err, context.DeadlineExceeded) {
		log.Println("Context deadline exceeded: creation user to db too long(register use case)")
		return nil, apperrors.NewDeadlineExceededError("user registration too long")
	}

	if errors.Is(err, context.Canceled) {
		log.Println("Context cancelled: creation user closed(register use case)")
		return nil, apperrors.NewCanceledError("user registration closed")
	}

	if err != nil {
		return nil, apperrors.NewInternalError("Unexpected internal error")
	}

	return &dto.CreateSessionOutputDto{
		SessionId: createdSession.Id(),
		Message:   "Session created",
	}, nil
}
