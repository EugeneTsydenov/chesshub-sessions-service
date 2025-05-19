package entity

import "time"

type Session struct {
	id         string
	userID     int64
	ipAddr     string
	deviceInfo string
	isActive   bool
	expiredAt  time.Time
	createdAt  time.Time
	updatedAt  time.Time
}

func (s *Session) ID() string {
	return s.id
}

func (s *Session) UserID() int64 {
	return s.userID
}

func (s *Session) IPAddr() string {
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

func (s *Session) UpdateIpAddr(newIpAddr string) {
	s.ipAddr = newIpAddr
}

func (s *Session) UpdateDeviceInfo(newDeviceInfo string) {
	s.deviceInfo = newDeviceInfo
}

func (s *Session) Activate() {
	s.isActive = true
}

func (s *Session) Deactivate() {
	s.isActive = false
}

func (s *Session) Refresh(t time.Time) {
	s.expiredAt = t
}

func (s *Session) Touch() {
	s.updatedAt = time.Now()
}
