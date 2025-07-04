package mapper

import (
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/dto"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/sessionfilter"
	sessionsproto "github.com/EugeneTsydenov/chesshub-sessions-service/internal/controllers/grpccontroller/genproto"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
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

func ToStopSessionInputDTO(req *sessionsproto.StopSessionRequest) *dto.StopSessionInputDTO {
	return &dto.StopSessionInputDTO{
		SessionID: req.SessionID,
	}
}

func ToStopSessionResponse(output *dto.StopSessionOutputDTO) *sessionsproto.StopSessionResponse {
	return &sessionsproto.StopSessionResponse{
		Message: output.Message,
	}
}

func ToListSessionInputDTO(req *sessionsproto.ListSessionsRequest) *dto.ListSessionsInputDTO {
	filter := req.Filter

	return &dto.ListSessionsInputDTO{
		Filter: &sessionfilter.SessionFilter{
			UserID:           filter.UserID,
			OnlyActive:       filter.OnlyActive,
			DeviceType:       toDeviceTypePtr(filter.DeviceType),
			DeviceName:       filter.DeviceName,
			AppType:          toAppTypePtr(filter.AppType),
			AppVersion:       filter.AppVersion,
			OS:               filter.Os,
			OSVersion:        filter.OsVersion,
			DeviceModel:      filter.DeviceModel,
			IPAddr:           filter.IpAddr,
			LastActiveBefore: toTimePtr(filter.LastActiveBefore),
			LastActiveAfter:  toTimePtr(filter.LastActiveAfter),
			UpdatedBefore:    toTimePtr(filter.UpdatedBefore),
			UpdatedAfter:     toTimePtr(filter.UpdatedAfter),
			CreatedBefore:    toTimePtr(filter.CreatedBefore),
			CreatedAfter:     toTimePtr(filter.CreatedAfter),
		},
	}
}

func toDeviceTypePtr(protoDeviceType *sessionsproto.DeviceType) *session.DeviceType {
	if protoDeviceType == nil {
		return nil
	}

	t := toDeviceType(*protoDeviceType)
	return &t
}

func toAppTypePtr(protoAppType *sessionsproto.AppType) *session.AppType {
	if protoAppType == nil {
		return nil
	}

	t := toAppType(*protoAppType)
	return &t
}

func toTimePtr(timestamp *timestamppb.Timestamp) *time.Time {
	if timestamp == nil {
		return nil
	}

	t := timestamp.AsTime()

	return &t
}

func ToListSessionsResponse(output *dto.ListSessionsOutputDTO) *sessionsproto.ListSessionsResponse {
	return &sessionsproto.ListSessionsResponse{
		Sessions: toListProtoSessions(output.Sessions),
		Count:    int32(output.Count),
		Message:  output.Message,
	}
}

func toListProtoSessions(sessions []*session.Session) []*sessionsproto.Session {
	var protoSessions []*sessionsproto.Session
	for _, s := range sessions {
		protoSessions = append(protoSessions, toProtoSession(s))
	}
	return protoSessions
}

func toProtoSession(session *session.Session) *sessionsproto.Session {
	return &sessionsproto.Session{
		Id:           session.ID().String(),
		UserID:       session.UserID(),
		DeviceType:   toProtoDeviceType(session.DeviceInfo().DeviceType()),
		DeviceName:   session.DeviceInfo().DeviceName(),
		AppType:      toProtoAppType(session.DeviceInfo().AppType()),
		AppVersion:   session.DeviceInfo().AppVersion(),
		Os:           session.DeviceInfo().OS(),
		OsVersion:    session.DeviceInfo().OSVersion(),
		DeviceModel:  session.DeviceInfo().DeviceModel(),
		IpAddr:       session.DeviceInfo().IPAddr(),
		City:         session.Location().City(),
		Country:      session.Location().Country(),
		IsActive:     session.IsActive(),
		Lifetime:     toProtoDuration(session.Lifetime()),
		LastActiveAt: toProtoTime(session.LastActiveAt()),
		UpdatedAt:    toProtoTime(session.UpdatedAt()),
		CreatedAt:    toProtoTime(session.CreatedAt()),
	}
}

func toProtoDeviceType(t session.DeviceType) sessionsproto.DeviceType {
	switch t {
	case session.DeviceTypeWeb:
		return sessionsproto.DeviceType_Web
	case session.DeviceTypeMobile:
		return sessionsproto.DeviceType_Mobile
	case session.DeviceTypeDesktop:
		return sessionsproto.DeviceType_Desktop
	case session.DeviceTypeTablet:
		return sessionsproto.DeviceType_Tablet
	default:
		return sessionsproto.DeviceType_Web // или sessionsproto.DeviceType(0), или специальное значение Unknown
	}
}

func toProtoAppType(t session.AppType) sessionsproto.AppType {
	switch t {
	case session.AppTypeChesshubWeb:
		return sessionsproto.AppType_ChesshubWeb
	case session.AppTypeChesshubMobile:
		return sessionsproto.AppType_ChesshubMobile
	case session.AppTypeChesshubDesktop:
		return sessionsproto.AppType_ChesshubDesktop
	case session.AppTypeChesshubTablet:
		return sessionsproto.AppType_ChesshubTablet
	default:
		return sessionsproto.AppType_ChesshubWeb
	}
}

func toProtoDuration(t time.Duration) *durationpb.Duration {
	return durationpb.New(t)
}

func toProtoTime(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}

func ToGetSessionInputDTO(req *sessionsproto.GetSessionRequest) *dto.GetSessionInputDTO {
	return &dto.GetSessionInputDTO{
		SessionID: req.SessionID,
	}
}

func ToGetSessionResponse(output *dto.GetSessionOutputDTO) *sessionsproto.GetSessionResponse {
	return &sessionsproto.GetSessionResponse{
		Session: toProtoSession(output.Session),
		Message: output.Message,
	}
}

func ToStopAllSessionsInputDTO(req *sessionsproto.StopAllSessionsRequest) *dto.StopAllSessionsInputDTO {
	return &dto.StopAllSessionsInputDTO{
		UserID: req.UserID,
	}
}

func ToStopAllSessionsResponse(output *dto.StopAllSessionsOutputDTO) *sessionsproto.StopAllSessionsResponse {
	return &sessionsproto.StopAllSessionsResponse{
		Message: output.Message,
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
