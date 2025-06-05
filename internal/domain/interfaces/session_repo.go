package interfaces

import (
	"context"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
	"github.com/google/uuid"
)

type SessionRepo interface {
	Create(ctx context.Context, session *session.Session) (*uuid.UUID, error)
}
