package dto

import "time"

type CreateSessionInputDTO struct {
	UserId     int64
	IpAddr     string
	DeviceInfo string
	ExpiredAt  time.Time
}

type CreateSessionOutputDTO struct {
	SessionId string
	Message   string
}
