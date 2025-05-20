package mapper

import (
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/dto"
	sessionsproto "github.com/EugeneTsydenov/chesshub-sessions-service/internal/controllers/grpccontroller/genproto"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToCreateSessionInputDTO(req *sessionsproto.CreateSessionRequest) *dto.CreateSessionInputDTO {
	return &dto.CreateSessionInputDTO{
		UserID:     req.UserId,
		IPAddr:     req.IpAddress,
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

func ToGetSessionByIDInputDTO(req *sessionsproto.GetSessionByIdRequest) *dto.GetSessionByIDInputDTO {
	return &dto.GetSessionByIDInputDTO{
		ID: req.Id,
	}
}

func ToGetSessionByIDResponse(output *dto.GetSessionByIDOutputDTO) *sessionsproto.GetSessionByIdResponse {
	return &sessionsproto.GetSessionByIdResponse{
		Session: toSessionData(output.Session),
		Message: output.Message,
	}
}

func ToGetSessionsInputDTO(req *sessionsproto.GetSessionsRequest) *dto.GetSessionsInputDTO {
	return &dto.GetSessionsInputDTO{
		FilterMap: req.FilterMap,
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
		Id:         session.ID(),
		IpAddress:  session.IPAddr(),
		DeviceInfo: session.DeviceInfo(),
		IsActive:   session.IsActive(),
		ExpiredAt:  timestamppb.New(session.ExpiredAt()),
		UpdatedAt:  timestamppb.New(session.UpdatedAt()),
		CreatedAt:  timestamppb.New(session.CreatedAt()),
	}
}

func ToUpdateSessionInputDTO(req *sessionsproto.UpdateSessionRequest) *dto.UpdateSessionInputDTO {
	return &dto.UpdateSessionInputDTO{
		SessionID: req.SessionId,
		FieldMap:  req.FieldMap,
	}
}

func ToUpdateSessionResponse(output *dto.UpdateSessionOutputDTO) *sessionsproto.UpdateSessionResponse {
	return &sessionsproto.UpdateSessionResponse{
		Session: toSessionData(output.Session),
		Message: output.Message,
	}
}

func ToDeactivateSessionInputDTO(req *sessionsproto.DeactivateSessionRequest) *dto.DeactivateSessionInputDTO {
	return &dto.DeactivateSessionInputDTO{
		SessionID: req.SessionId,
	}
}

func ToDeactivateSessionResponse(output *dto.DeactivateSessionOutputDTO) *sessionsproto.DeactivateSessionResponse {
	return &sessionsproto.DeactivateSessionResponse{
		Success: output.Success,
		Message: output.Message,
	}
}
