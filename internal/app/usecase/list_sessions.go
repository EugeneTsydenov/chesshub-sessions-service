package usecase

import (
	"context"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/dto"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/sessionfilter"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/interfaces"
)

type (
	ListSessions UseCase[*dto.ListSessionsInputDTO, *dto.ListSessionsOutputDTO]

	listSessions struct {
		filterBuilder sessionfilter.Builder
		sessionRepo   interfaces.SessionRepo
	}
)

var _ ListSessions = new(listSessions)

func NewListSessions(filterBuilder sessionfilter.Builder, sessionRepo interfaces.SessionRepo) ListSessions {
	return &listSessions{
		filterBuilder: filterBuilder,
		sessionRepo:   sessionRepo,
	}
}

func (uc listSessions) Execute(ctx context.Context, input *dto.ListSessionsInputDTO) (*dto.ListSessionsOutputDTO, error) {
	criteria := uc.filterBuilder.BuildCriteria(input.Filter)

	sessions, err := uc.sessionRepo.Find(ctx, criteria)
	if err != nil {
		return nil, err
	}

	return &dto.ListSessionsOutputDTO{
		Sessions: sessions,
		Count:    len(sessions),
		Message:  "Sessions received",
	}, nil
}
