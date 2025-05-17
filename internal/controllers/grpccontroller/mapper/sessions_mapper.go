package mapper

import (
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/dto"
	sessionsproto "github.com/EugeneTsydenov/chesshub-sessions-service/internal/controllers/grpccontroller/genproto"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func ToCreateSessionInputDTO(req *sessionsproto.CreateSessionRequest) *dto.CreateSessionInputDTO {
	return &dto.CreateSessionInputDTO{
		UserId:     req.UserId,
		IpAddr:     req.IpAddress,
		DeviceInfo: req.DeviceInfo,
		ExpiredAt:  req.ExpiredAt.AsTime(),
	}
}

func ToCreateSessionResponse(output *dto.CreateSessionOutputDTO) *sessionsproto.CreateSessionResponse {
	return &sessionsproto.CreateSessionResponse{
		Session: toSessionData(output.Session),
		Message: output.Message,
	}
}

func ToGetSessionByIDInputDTO(req *sessionsproto.GetSessionByIdRequest) *dto.GetSessionByIdInputDTO {
	return &dto.GetSessionByIdInputDTO{
		ID: req.Id,
	}
}

func ToGetSessionByIDResponse(output *dto.GetSessionByIdOutputDTO) *sessionsproto.GetSessionByIdResponse {
	return &sessionsproto.GetSessionByIdResponse{
		Session: toSessionData(output.Session),
		Message: output.Message,
	}
}

func ToGetSessionsInputDTO(req *sessionsproto.GetSessionsRequest) *dto.GetSessionsInputDTO {
	var expiredBefore, expiredAfter time.Time
	if req.ExpiredBefore != nil {
		expiredBefore = req.ExpiredBefore.AsTime()
	}
	if req.ExpiredAfter != nil {
		expiredAfter = req.ExpiredAfter.AsTime()
	}

	return &dto.GetSessionsInputDTO{
		UserId:        req.UserId,
		IpAddr:        req.IpAddress,
		DeviceInfo:    req.DeviceInfo,
		IsActive:      req.IsActive,
		ExpiredBefore: expiredBefore,
		ExpiredAfter:  expiredAfter,
	}
}

func ToGetSessionsResponse(output *dto.GetSessionsOutputDTO) *sessionsproto.GetSessionsResponse {
	converted := make([]*sessionsproto.SessionData, len(output.Sessions), cap(output.Sessions))
	for i, v := range output.Sessions {
		converted[i] = toSessionData(v)
	}

	return &sessionsproto.GetSessionsResponse{
		Sessions: converted,
		Message:  output.Message,
	}
}

func toSessionData(session *entity.Session) *sessionsproto.SessionData {
	return &sessionsproto.SessionData{
		Id:         session.Id(),
		IpAddress:  session.IpAddr(),
		DeviceInfo: session.DeviceInfo(),
		IsActive:   session.IsActive(),
		ExpiredAt:  timestamppb.New(session.ExpiredAt()),
		UpdatedAt:  timestamppb.New(session.UpdatedAt()),
		CreatedAt:  timestamppb.New(session.CreatedAt()),
	}
}

func ToUpdateSessionInputDTO(req *sessionsproto.UpdateSessionRequest) *dto.UpdateSessionInputDTO {
	return &dto.UpdateSessionInputDTO{
		UserId:     req.UserId,
		IpAddr:     req.IpAddress,
		DeviceInfo: req.DeviceInfo,
		IsActive:   req.IsActive,
		ExpiredAt:  req.ExpiredAt.AsTime(),
	}
}

func ToUpdateSessionResponse(output *dto.UpdateSessionOutputDTO) *sessionsproto.UpdateSessionResponse {
	return &sessionsproto.UpdateSessionResponse{
		Session: toSessionData(output.Session),
	}
}
