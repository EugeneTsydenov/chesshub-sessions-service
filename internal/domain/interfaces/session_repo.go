package interfaces

import (
	"context"
	"time"

	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
	"github.com/google/uuid"
)

type (
	SessionRepo interface {
		Create(ctx context.Context, session *session.Session) (*session.Session, error)
		GetByID(ctx context.Context, id uuid.UUID) (*session.Session, error)
		Find(ctx context.Context, criteria *session.Criteria) ([]*session.Session, error)
		Update(ctx context.Context, session *session.Session) (*session.Session, error)
	}

	SessionCache interface {
		HSet(ctx context.Context, s *session.Session) error
		HGet(ctx context.Context, id uuid.UUID) (*session.Session, error)
		Del(ctx context.Context, id uuid.UUID) error
		Exists(ctx context.Context, id uuid.UUID) (bool, error)
		ExtendTTL(ctx context.Context, id uuid.UUID, ttl time.Duration) error
	}
)
