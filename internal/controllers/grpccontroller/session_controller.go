package grpccontroller

import (
	"context"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/usecase"
	sessionsproto "github.com/EugeneTsydenov/chesshub-sessions-service/internal/controllers/grpccontroller/genproto"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/controllers/grpccontroller/mapper"
	"time"
)

type SessionController struct {
	sessionsproto.UnimplementedSessionsServiceServer
	createSessionUseCase  usecase.CreateSessionUseCase
	getSessionByIdUseCase usecase.GetSessionByIdUseCase
	getSessionsUseCase    usecase.GetSessionsUseCase
}

func NewSessionController(
	createSessionUseCase usecase.CreateSessionUseCase,
	getSessionByIdUseCase usecase.GetSessionByIdUseCase,
	getSessionsUseCase usecase.GetSessionsUseCase,
) *SessionController {
	return &SessionController{
		createSessionUseCase:  createSessionUseCase,
		getSessionByIdUseCase: getSessionByIdUseCase,
		getSessionsUseCase:    getSessionsUseCase,
	}
}

func (c *SessionController) CreateSession(ctx context.Context, req *sessionsproto.CreateSessionRequest) (*sessionsproto.CreateSessionResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	r, err := c.createSessionUseCase.Execute(ctx, mapper.ToCreateSessionInputDTO(req))
	if err != nil {
		return nil, err
	}

	return mapper.ToCreateSessionResponse(r), nil
}

func (c *SessionController) GetSessionById(ctx context.Context, req *sessionsproto.GetSessionByIdRequest) (*sessionsproto.GetSessionByIdResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	r, err := c.getSessionByIdUseCase.Execute(ctx, mapper.ToGetSessionByIDInputDTO(req))
	if err != nil {
		return nil, err
	}

	return mapper.ToGetSessionByIDResponse(r), nil
}

func (c *SessionController) GetSessions(ctx context.Context, req *sessionsproto.GetSessionsRequest) (*sessionsproto.GetSessionsResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	r, err := c.getSessionsUseCase.Execute(ctx, mapper.ToGetSessionsInputDTO(req))
	if err != nil {
		return nil, err
	}

	return mapper.ToGetSessionsResponse(r), nil
}

//func (c *SessionController) UpdateSession(ctx context.Context, req *sessionsproto.UpdateSessionRequest) (*sessionsproto.UpdateSessionResponse, error) {
//	if err := req.ValidateAll(); err != nil {
//		return nil, err
//	}
//
//	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
//	defer cancel()
//
//	r, err := c.updateSessionUseCase.Execute(ctx, mapper.ToUpdateSessionInputDTO(req))
//	if err != nil {
//		return nil, err
//	}
//
//	return mapper.ToUpdateSessionResponse(r), nil
//}
