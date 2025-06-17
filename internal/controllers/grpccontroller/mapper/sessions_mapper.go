package mapper

import (
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/dto"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/controllers/grpccontroller/genproto"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
)

func ToStartSessionInputDTO(req *sessionsproto.StartSessionRequest) *dto.StartSessionInputDTO {
	return &dto.StartSessionInputDTO{
		UserID:     req.UserId,
		DeviceInfo: toDeviceInfo(req.DeviceInfo),
	}
}

func toDeviceInfo(input *sessionsproto.DeviceInfo) *session.DeviceInfo {
	return session.NewDeviceInfo(
		toDeviceType(input.DeviceType),
		input.DeviceName,
		toAppType(input.AppType),
		input.AppVersion,
		input.Os,
		input.OsVersion,
		input.DeviceModel,
		input.IpAddr,
	)
}

func toDeviceType(input sessionsproto.DeviceType) session.DeviceType {
	return session.DeviceType(input)
}

func toAppType(input sessionsproto.AppType) session.AppType {
	return session.AppType(input)
}

func ToStartSessionResponse(output *dto.StartSessionOutputDTO) *sessionsproto.StartSessionResponse {
	return &sessionsproto.StartSessionResponse{
		SessionID: output.SessionID.String(),
		Message:   output.Message,
	}
}

//func ToGetSessionByIDInputDTO(req *sessionsproto.GetSessionByIdRequest) *dto.GetSessionByIDInputDTO {
//	return &dto.GetSessionByIDInputDTO{
//		ID: req.Id,
//	}
//}
//
//func ToGetSessionByIDResponse(output *dto.GetSessionByIDOutputDTO) *sessionsproto.GetSessionByIdResponse {
//	return &sessionsproto.GetSessionByIdResponse{
//		Session: toSessionData(output.Session),
//		Message: output.Message,
//	}
//}
//
//func ToGetSessionsInputDTO(req *sessionsproto.GetSessionsRequest) *dto.GetSessionsInputDTO {
//	return &dto.GetSessionsInputDTO{
//		FilterMap: req.FilterMap,
//	}
//}
//
//func ToGetSessionsResponse(output *dto.GetSessionsOutputDTO) *sessionsproto.GetSessionsResponse {
//	converted := make([]*sessionsproto.SessionData, len(output.Sessions), cap(output.Sessions))
//	for i, v := range output.Sessions {
//		converted[i] = toSessionData(v)
//	}
//
//	return &sessionsproto.GetSessionsResponse{
//		Sessions: converted,
//		Message:  output.Message,
//	}
//}
//
//func toSessionData(session *session.Session) *sessionsproto.SessionData {
//	return &sessionsproto.SessionData{
//		Id:         session.ID(),
//		IpAddress:  session.IPAddr(),
//		DeviceInfo: session.DeviceInfo(),
//		IsActive:   session.IsActive(),
//		ExpiredAt:  timestamppb.New(session.ExpiredAt()),
//		UpdatedAt:  timestamppb.New(session.UpdatedAt()),
//		CreatedAt:  timestamppb.New(session.CreatedAt()),
//	}
//}
//
//func ToUpdateSessionInputDTO(req *sessionsproto.UpdateSessionRequest) *dto.UpdateSessionInputDTO {
//	return &dto.UpdateSessionInputDTO{
//		SessionID: req.SessionId,
//		FieldMap:  req.FieldMap,
//	}
//}
//
//func ToUpdateSessionResponse(output *dto.UpdateSessionOutputDTO) *sessionsproto.UpdateSessionResponse {
//	return &sessionsproto.UpdateSessionResponse{
//		Session: toSessionData(output.Session),
//		Message: output.Message,
//	}
//}
//
//func ToDeactivateSessionInputDTO(req *sessionsproto.DeactivateSessionRequest) *dto.DeactivateSessionInputDTO {
//	return &dto.DeactivateSessionInputDTO{
//		SessionID: req.SessionId,
//	}
//}
//
//func ToDeactivateSessionResponse(output *dto.DeactivateSessionOutputDTO) *sessionsproto.DeactivateSessionResponse {
//	return &sessionsproto.DeactivateSessionResponse{
//		Success: output.Success,
//		Message: output.Message,
//	}
//}
