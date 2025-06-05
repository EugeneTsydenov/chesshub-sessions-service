package session

import (
	"github.com/google/uuid"
	"time"
)

type Builder struct {
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

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) WithID(id uuid.UUID) *Builder {
	b.id = id
	return b
}

func (b *Builder) WithUserID(userID int64) *Builder {
	b.userID = userID
	return b
}

func (b *Builder) WithDeviceInfo(deviceInfo *DeviceInfo) *Builder {
	b.deviceInfo = deviceInfo
	return b
}

func (b *Builder) WithLocation(location *Location) *Builder {
	b.location = location
	return b
}

func (b *Builder) WithIsActive(isActive bool) *Builder {
	b.isActive = isActive
	return b
}

func (b *Builder) WithLifetime(lifetime time.Duration) *Builder {
	b.lifetime = lifetime
	return b
}

func (b *Builder) WithLastActiveAt(lastActiveAt time.Time) *Builder {
	b.lastActiveAt = lastActiveAt
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
	return newSession(b)
}
