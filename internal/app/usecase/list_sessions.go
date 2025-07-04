package usecase

import (
	"context"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/dto"
	apperrors "github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/errors"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/sessionfilter"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/interfaces"
)

type (
	ListSessions UseCase[*dto.ListSessionsInputDTO, *dto.ListSessionsOutputDTO]

	listSessions struct {
		filterBuilder     sessionfilter.Builder
		cachedSessionRepo interfaces.SessionRepo
	}
)

var _ ListSessions = new(listSessions)

func NewListSessions(filterBuilder sessionfilter.Builder, cachedRepo interfaces.SessionRepo) ListSessions {
	return &listSessions{
		filterBuilder:     filterBuilder,
		cachedSessionRepo: cachedRepo,
	}
}

func (uc listSessions) Execute(ctx context.Context, input *dto.ListSessionsInputDTO) (*dto.ListSessionsOutputDTO, error) {
	criteria := uc.filterBuilder.BuildCriteria(input.Filter)

	sessions, err := uc.cachedSessionRepo.Find(ctx, criteria)
	if err != nil {
		return nil, apperrors.FromDomainError(err)
	}

	return &dto.ListSessionsOutputDTO{
		Sessions: sessions,
		Count:    len(sessions),
		Message:  "Sessions received",
	}, nil
}
