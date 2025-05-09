package dto

import "time"

type CreateSessionInputDto struct {
	UserId     int64
	IpAddr     string
	DeviceInfo string
	ExpiredAt  time.Time
}
