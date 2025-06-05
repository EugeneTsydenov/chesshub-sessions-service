package session

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrGenerateID = errors.New("error generating session id")
)

type Session struct {
	id           uuid.UUID
	userID       int64
	deviceInfo   *DeviceInfo
	location     *Location
	isActive     bool
	lifetime     time.Duration
	lastActiveAt time.Time
	createdAt    time.Time
	updatedAt    time.Time
}

func newSession(b *Builder) *Session {
	return &Session{
		id:           b.id,
		userID:       b.userID,
		deviceInfo:   b.deviceInfo,
		location:     b.location,
		isActive:     b.isActive,
		lastActiveAt: b.lastActiveAt,
		createdAt:    b.createdAt,
		updatedAt:    b.updatedAt,
	}
}

func (s *Session) ID() uuid.UUID {
	return s.id
}

func (s *Session) UserID() int64 {
	return s.userID
}

func (s *Session) DeviceInfo() *DeviceInfo {
	return s.deviceInfo
}

func (s *Session) Location() *Location {
	return s.location
}

func (s *Session) IsActive() bool {
	return s.isActive
}

func (s *Session) Lifetime() time.Duration {
	return s.lifetime
}

func (s *Session) LastActiveAt() time.Time {
	return s.lastActiveAt
}

func (s *Session) CreatedAt() time.Time {
	return s.createdAt
}

func (s *Session) UpdatedAt() time.Time {
	return s.updatedAt
}

func (s *Session) Initialize() error {
	if err := s.GenerateID(); err != nil {
		return err
	}
	s.Activate()
	s.RefreshLastActiveAt()
	return nil
}

func (s *Session) GenerateID() error {
	id, err := uuid.NewRandom()
	if err != nil {
		return ErrGenerateID
	}

	s.id = id

	return nil
}

func (s *Session) Activate() {
	s.isActive = true
}

func (s *Session) RefreshLastActiveAt() {
	s.lastActiveAt = time.Now()
}

func (s *Session) UpdateLocation(location *Location) {
	s.location = location
}

//func (s *Session) UpdateIpAddr(newIpAddr string) {
//	s.ipAddr = newIpAddr
//}
//
//func (s *Session) UpdateDeviceInfo(newDeviceInfo string) {
//	s.deviceInfo = newDeviceInfo
//}
//
//func (s *Session) Activate() {
//	s.isActive = true
//}
//
//func (s *Session) Deactivate() {
//	s.isActive = false
//}
//
//func (s *Session) Extend(d time.Duration) {
//	s.expiredAt = time.Now().Add(d)
//}
//
//func (s *Session) Touch() {
//	s.updatedAt = time.Now()
//}
