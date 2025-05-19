package usecase

import (
	"context"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/dto"
	apperrors "github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/errors"
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
	session, err := u.sessionsRepo.GetByID(ctx, input.SessionID)
	if err != nil {
		return nil, apperrors.NewNotFoundError("session not found", err)
	}

	if input.Fields.IpAddr != nil {
		session.UpdateIpAddr(*input.Fields.IpAddr)
	}

	if input.Fields.DeviceInfo != nil {
		session.UpdateDeviceInfo(*input.Fields.DeviceInfo)
	}

	if input.Fields.IsActive != nil {
		if *input.Fields.IsActive {
			session.Activate()
		} else {
			session.Deactivate()
		}
	}
	return nil, nil
}
