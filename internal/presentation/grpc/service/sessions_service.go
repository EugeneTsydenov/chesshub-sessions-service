package service

import (
	"context"

	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/dto"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/usecase"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/presentation/grpc/generated/sessions"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/presentation/grpc/grpcerrors"
)

type SessionsService struct {
	sessions.UnimplementedSessionsServiceServer
	createSessionUseCase usecase.CreateSessionUseCase
}

func NewSessionsService(createSessionUseCase usecase.CreateSessionUseCase) *SessionsService {
	return &SessionsService{
		createSessionUseCase: createSessionUseCase,
	}
}

func (s *SessionsService) CreateSession(ctx context.Context, req *sessions.CreateSessionRequest) (*sessions.CreateSessionResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	inputDto := &dto.CreateSessionInputDto{
		UserId:     req.UserId,
		IpAddr:     req.IpAddress,
		DeviceInfo: req.DeviceInfo,
		ExpiredAt:  req.ExpiredAt.AsTime(),
	}

	r, err := s.createSessionUseCase.Execute(ctx, inputDto)
	if err != nil {
		return nil, grpcerrors.MapAppErrorToGrpcError(err)
	}

	return &sessions.CreateSessionResponse{
		SessionId: r.SessionId,
		Message:   r.Message,
	}, nil
}
