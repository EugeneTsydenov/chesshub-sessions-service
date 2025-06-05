package interfaces

import "github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"

type SessionService interface {
	EnrichLocation(s *session.Session)
}
