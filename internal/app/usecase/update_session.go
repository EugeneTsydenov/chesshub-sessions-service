package usecase

import (
	"context"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/dto"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/port"
)

type UpdateSessionUseCase interface {
	UseCase[*dto.UpdateSessionInputDTO, *dto.UpdateSessionOutputDTO]
}

type UpdateSessionUseCaseImpl struct {
	sessionsRepo port.SessionsRepo
}

var _ UpdateSessionUseCase = new(UpdateSessionUseCaseImpl)

func NewUpdateSessionUseCase(sessionsRepo port.SessionsRepo) *UpdateSessionUseCaseImpl {
	return &UpdateSessionUseCaseImpl{
		sessionsRepo: sessionsRepo,
	}
}

func (u UpdateSessionUseCaseImpl) Execute(ctx context.Context, input *dto.UpdateSessionInputDTO) (*dto.UpdateSessionOutputDTO, error) {
	//TODO implement me
	panic("implement me")
}
