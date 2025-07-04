package usecase

import (
	"context"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/dto"
	apperrors "github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/errors"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/sessionfilter"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/interfaces"
	"sync"
)

type (
	StopAllSessions UseCase[*dto.StopAllSessionsInputDTO, *dto.StopAllSessionsOutputDTO]

	stopAllSessions struct {
		filterBuilder     sessionfilter.Builder
		cachedSessionRepo interfaces.SessionRepo
		sessionService    interfaces.SessionService
	}
)

func NewStopAllSessions(filterBuilder sessionfilter.Builder, cachedSessionRepo interfaces.SessionRepo, sessionService interfaces.SessionService) StopAllSessions {
	return &stopAllSessions{
		filterBuilder:     filterBuilder,
		cachedSessionRepo: cachedSessionRepo,
		sessionService:    sessionService,
	}
}

func (uc *stopAllSessions) Execute(ctx context.Context, input *dto.StopAllSessionsInputDTO) (*dto.StopAllSessionsOutputDTO, error) {
	filter := &sessionfilter.SessionFilter{
		UserID: &input.UserID,
	}

	criteria := uc.filterBuilder.BuildCriteria(filter)

	sessions, err := uc.cachedSessionRepo.Find(ctx, criteria)
	if err != nil {
		return nil, apperrors.NewInternalError(err.Error())
	}

	var wg sync.WaitGroup

	for _, s := range sessions {
		wg.Add(1)

		go func(s *session.Session) {
			defer wg.Done()
			_ = uc.sessionService.DeactivateSession(ctx, s)
		}(s)
	}

	wg.Wait()

	return &dto.StopAllSessionsOutputDTO{
		Message: "All sessions have been deactivated",
	}, nil
}
