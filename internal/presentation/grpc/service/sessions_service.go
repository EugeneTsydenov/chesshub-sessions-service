package service

import (
	"context"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/usecase"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/presentation/grpc/generated/sessions"
	"log"
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

func (s *SessionsService) CreateSession(_ context.Context, _ *sessions.CreateSessionRequest) (*sessions.CreateSessionResponse, error) {
	log.Print("CreateSession called")

	_, _ = s.createSessionUseCase.Execute(context.Background(), nil)

	return &sessions.CreateSessionResponse{
		SessionId: "10",
		Message:   "SUCCESS",
	}, nil
}
