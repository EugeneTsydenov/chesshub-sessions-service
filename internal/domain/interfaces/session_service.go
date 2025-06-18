package interfaces

import (
	"context"

	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
)

type SessionService interface {
	EnrichLocation(s *session.Session)
	DeactivateSession(ctx context.Context, s *session.Session) error
}
