package usecase

import (
	"context"
	"errors"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/dto"
	apperrors "github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/errors"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/port"
)

type (
	GetSessionByIdUseCase interface {
		UseCase[*dto.GetSessionByIDInputDTO, *dto.GetSessionByIDOutputDTO]
	}

	GetSessionByIdUseCaseImpl struct {
		sessionsRepo port.SessionsRepo
	}
)

var _ GetSessionByIdUseCase = new(GetSessionByIdUseCaseImpl)

func NewGetSessionByIdUseCase(sessionsRepo port.SessionsRepo) GetSessionByIdUseCase {
	return &GetSessionByIdUseCaseImpl{
		sessionsRepo: sessionsRepo,
	}
}

func (u *GetSessionByIdUseCaseImpl) Execute(ctx context.Context, input *dto.GetSessionByIDInputDTO) (*dto.GetSessionByIDOutputDTO, error) {
	session, err := u.sessionsRepo.GetByID(ctx, input.ID)

	if errors.Is(err, context.DeadlineExceeded) {
		return nil, apperrors.NewDeadlineExceededError("getting session by id too long", err)
	}

	if errors.Is(err, context.Canceled) {
		return nil, apperrors.NewCanceledError("getting session by id closed", err)
	}

	if err != nil {
		return nil, apperrors.NewInternalError("something went wrong", err)
	}

	return &dto.GetSessionByIDOutputDTO{
		Session: session,
		Message: "Success",
	}, nil
}
