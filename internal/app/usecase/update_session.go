package usecase

import (
	"context"
	"errors"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/dto"
	apperrors "github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/errors"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/port"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/specs/sessionspec"
	"log"
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
	log.Print(input.FieldMap)
	spec := sessionspec.NewSessionUpdateSpec(input.SessionID, input.FieldMap)

	updatedSession, err := u.sessionsRepo.Update(ctx, spec)

	if errors.Is(err, context.DeadlineExceeded) {
		return nil, apperrors.NewDeadlineExceededError("update session too long", err)
	}

	if errors.Is(err, context.Canceled) {
		return nil, apperrors.NewCanceledError("update session closed", err)
	}

	if err != nil {
		return nil, err
	}

	return &dto.UpdateSessionOutputDTO{
		Session: updatedSession,
		Message: "Session updated",
	}, nil
}
