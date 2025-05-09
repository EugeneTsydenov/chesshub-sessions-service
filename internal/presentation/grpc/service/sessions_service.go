package service

import (
	"context"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/presentation/grpc/generated/sessions"
)

type SessionsService struct {
	sessions.UnimplementedSessionsServiceServer
}

func NewSessionsService() *SessionsService {
	return &SessionsService{}
}

func (s *SessionsService) CreateSession(_ context.Context, _ *sessions.CreateSessionRequest) (*sessions.CreateSessionResponse, error) {
	return &sessions.CreateSessionResponse{
		SessionId: "10",
		Message:   "SUCCESS",
	}, nil
}
