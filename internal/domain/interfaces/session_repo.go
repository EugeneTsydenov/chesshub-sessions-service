package interfaces

import (
	"context"

	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
	"github.com/google/uuid"
)

type SessionRepo interface {
	Create(ctx context.Context, session *session.Session) (*uuid.UUID, error)
	GetByID(ctx context.Context, id uuid.UUID) (*session.Session, error)
	Update(ctx context.Context, session *session.Session) (*uuid.UUID, error)
}
