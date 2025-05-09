package session

import "time"

type Builder struct {
	id         string
	userId     int64
	ipAddr     string
	deviceInfo string
	isActive   bool
	expiredAt  time.Time
	createdAt  time.Time
	updatedAt  time.Time
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) WithId(id string) *Builder {
	b.id = id
	return b
}

func (b *Builder) WithUserId(userId int64) *Builder {
	b.userId = userId
	return b
}

func (b *Builder) WithIpAddr(ipAddr string) *Builder {
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
		userId:     b.userId,
		ipAddr:     b.ipAddr,
		deviceInfo: b.deviceInfo,
		expiredAt:  b.expiredAt,
		createdAt:  b.createdAt,
		updatedAt:  b.updatedAt,
	}
}
