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

	if input.IpAddr != nil {
		session.UpdateIpAddr(*input.IpAddr)
	}

	if input.DeviceInfo != nil {
		session.UpdateDeviceInfo(*input.DeviceInfo)
	}

	if input.IsActive != nil || *input.IsActive {
		session.Activate()
	} else {
		session.Deactivate()
	}

	if !input.ExpiredAt.IsZero() {
		session.Refresh(input.ExpiredAt)
	}

	session.Touch()

	updatedSession, err := u.sessionsRepo.Update(ctx, session)

	return &dto.UpdateSessionOutputDTO{
		Session: updatedSession,
		Message: "Session updated",
	}, nil
}
