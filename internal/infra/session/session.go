package session

import "time"

type Session struct {
	id         string
	userId     int64
	ipAddr     string
	deviceInfo string
	isActive   bool
	expiredAt  time.Time
	createdAt  time.Time
	updatedAt  time.Time
}

func (s *Session) Id() string {
	return s.id
}

func (s *Session) UserId() int64 {
	return s.userId
}

func (s *Session) IpAddr() string {
	return s.ipAddr
}

func (s *Session) DeviceInfo() string {
	return s.deviceInfo
}

func (s *Session) IsActive() bool {
	return s.isActive
}

func (s *Session) ExpiredAt() time.Time {
	return s.expiredAt
}

func (s *Session) CreatedAt() time.Time {
	return s.createdAt
}

func (s *Session) UpdatedAt() time.Time {
	return s.updatedAt
}
