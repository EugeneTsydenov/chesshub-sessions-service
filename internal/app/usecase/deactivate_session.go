package usecase

import (
	"context"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/dto"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/errors"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/port"
)

type (
	DeactivateSessionUseCase interface {
		UseCase[*dto.DeactivateSessionInputDTO, *dto.DeactivateSessionOutputDTO]
	}

	DeactivateSessionUseCaseImpl struct {
		sessionsRepo port.SessionsRepo
	}
)

var _ DeactivateSessionUseCase = new(DeactivateSessionUseCaseImpl)

func NewDeactivateSessionUseCase(sessionsRepo port.SessionsRepo) *DeactivateSessionUseCaseImpl {
	return &DeactivateSessionUseCaseImpl{
		sessionsRepo: sessionsRepo,
	}
}

func (u DeactivateSessionUseCaseImpl) Execute(ctx context.Context, input *dto.DeactivateSessionInputDTO) (*dto.DeactivateSessionOutputDTO, error) {
	session, err := u.sessionsRepo.GetByID(ctx, input.SessionID)
	if err != nil {
		return nil, errors.NewInternalError("Something went wrong!", err)
	}

	if session == nil {
		return nil, errors.NewNotFoundError("Session not found!", nil)
	}

	session.Deactivate()
	session.Touch()

	_, err = u.sessionsRepo.Update(ctx, session)
	if err != nil {
		return nil, errors.NewInternalError("Something went wrong!", err)
	}

	return &dto.DeactivateSessionOutputDTO{
		Success: true,
		Message: "Session successfully deactivated!",
	}, nil
}
