package grpccontroller

import (
	"context"
	"time"

	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/usecase"
	sessionsproto "github.com/EugeneTsydenov/chesshub-sessions-service/internal/controllers/grpccontroller/genproto"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/controllers/grpccontroller/mapper"
)

type SessionController struct {
	sessionsproto.UnimplementedSessionsServiceServer
	startSessionUseCase usecase.StartSession
	stopSessionUseCase  usecase.StopSession
	listSessionsUseCase usecase.ListSessions
}

func NewSessionController(
	startSessionUseCase usecase.StartSession,
	stopSessionUseCase usecase.StopSession,
	listSessionsUseCase usecase.ListSessions,
) *SessionController {
	return &SessionController{
		startSessionUseCase: startSessionUseCase,
		stopSessionUseCase:  stopSessionUseCase,
		listSessionsUseCase: listSessionsUseCase,
	}
}

func (c *SessionController) StartSession(ctx context.Context, req *sessionsproto.StartSessionRequest) (*sessionsproto.StartSessionResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	r, err := c.startSessionUseCase.Execute(ctx, mapper.ToStartSessionInputDTO(req))
	if err != nil {
		return nil, err
	}

	return mapper.ToStartSessionResponse(r), nil
}

func (c *SessionController) StopSession(ctx context.Context, req *sessionsproto.StopSessionRequest) (*sessionsproto.StopSessionResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	r, err := c.stopSessionUseCase.Execute(ctx, mapper.ToStopSessionInputDTO(req))
	if err != nil {
		return nil, err
	}

	return mapper.ToStopSessionResponse(r), nil
}

func (c *SessionController) ListSessions(ctx context.Context, req *sessionsproto.ListSessionsRequest) (*sessionsproto.ListSessionsResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	r, err := c.listSessionsUseCase.Execute(ctx, mapper.ToListSessionInputDTO(req))
	if err != nil {
		return nil, err
	}

	return mapper.ToListSessionsResponse(r), nil
}

//
//func (c *SessionController) GetSessionById(ctx context.Context, req *sessionsproto.GetSessionByIdRequest) (*sessionsproto.GetSessionByIdResponse, error) {
//	if err := req.ValidateAll(); err != nil {
//		return nil, err
//	}
//
//	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
//	defer cancel()
//
//	r, err := c.getSessionByIdUseCase.Execute(ctx, mapper.ToGetSessionByIDInputDTO(req))
//	if err != nil {
//		return nil, err
//	}
//
//	return mapper.ToGetSessionByIDResponse(r), nil
//}
//
//func (c *SessionController) GetSessions(ctx context.Context, req *sessionsproto.GetSessionsRequest) (*sessionsproto.GetSessionsResponse, error) {
//	if err := req.ValidateAll(); err != nil {
//		return nil, err
//	}
//
//	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
//	defer cancel()
//
//	r, err := c.getSessionsUseCase.Execute(ctx, mapper.ToGetSessionsInputDTO(req))
//	if err != nil {
//		return nil, err
//	}
//
//	return mapper.ToGetSessionsResponse(r), nil
//}
//
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
//
//func (c *SessionController) DeactivateSession(ctx context.Context, req *sessionsproto.DeactivateSessionRequest) (*sessionsproto.DeactivateSessionResponse, error) {
//	if err := req.ValidateAll(); err != nil {
//		return nil, err
//	}
//
//	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
//	defer cancel()
//
//	r, err := c.deactivateSessionUseCase.Execute(ctx, mapper.ToDeactivateSessionInputDTO(req))
//	if err != nil {
//		return nil, err
//	}
//
//	return mapper.ToDeactivateSessionResponse(r), nil
//}
