package usecase

import (
	"context"
	"log"

	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/dto"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/infra/session/data/repo"
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

func (u *CreateSessionUseCaseImpl) Execute(_ context.Context, input *dto.CreateSessionInputDto) (*dto.CreateSessionOutputDto, error) {
	log.Print("CreateSessionUseCase Execute called")

	_, _ = u.sessionsRepo.Create(context.Background(), nil)

	return &dto.CreateSessionOutputDto{
		SessionId: "sakhjdkajshfkasdjfhsdkj",
		Message:   "Session created",
	}, nil
}
