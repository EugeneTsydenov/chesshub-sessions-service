package entity

import (
	"time"
)

type Builder struct {
	id         string
	userID     int64
	ipAddr     string
	deviceInfo string
	isActive   bool
	expiredAt  time.Time
	createdAt  time.Time
	updatedAt  time.Time
}

func NewSessionBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) WithID(id string) *Builder {
	b.id = id
	return b
}

func (b *Builder) WithUserID(userID int64) *Builder {
	b.userID = userID
	return b
}

func (b *Builder) WithIPAddr(ipAddr string) *Builder {
	b.ipAddr = ipAddr
	return b
}

func (b *Builder) WithDeviceInfo(deviceInfo string) *Builder {
	b.deviceInfo = deviceInfo
	return b
}

func (b *Builder) WithIsActive(isActive bool) *Builder {
	b.isActive = isActive
	return b
}

func (b *Builder) WithExpiredAt(expiredAt time.Time) *Builder {
	b.expiredAt = expiredAt
	return b
}

func (b *Builder) WithCreatedAt(createdAt time.Time) *Builder {
	b.createdAt = createdAt
	return b
}

func (b *Builder) WithUpdatedAt(updatedAt time.Time) *Builder {
	b.updatedAt = updatedAt
	return b
}

func (b *Builder) Build() *Session {
	return &Session{
		id:         b.id,
		userID:     b.userID,
		ipAddr:     b.ipAddr,
		deviceInfo: b.deviceInfo,
		isActive:   b.isActive,
		expiredAt:  b.expiredAt,
		createdAt:  b.createdAt,
		updatedAt:  b.updatedAt,
	}
}
