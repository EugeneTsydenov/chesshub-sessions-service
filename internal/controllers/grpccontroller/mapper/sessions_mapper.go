package mapper

import (
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/app/dto"
	sessionsproto "github.com/EugeneTsydenov/chesshub-sessions-service/internal/controllers/grpccontroller/genproto"
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
		SessionId: output.SessionId,
		Message:   output.Message,
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
	converted := make([]*sessionsproto.SessionResponse, len(output.Sessions), cap(output.Sessions))
	for i, v := range output.Sessions {
		converted[i] = &sessionsproto.SessionResponse{
			Id:         v.Id(),
			IpAddress:  v.IpAddr(),
			DeviceInfo: v.DeviceInfo(),
			IsActive:   v.IsActive(),
			ExpiredAt:  timestamppb.New(v.ExpiredAt()),
			UpdatedAt:  timestamppb.New(v.UpdatedAt()),
			CreatedAt:  timestamppb.New(v.CreatedAt()),
		}
	}

	return &sessionsproto.GetSessionsResponse{
		Sessions: converted,
		Message:  output.Message,
	}
}
