package services

import (
	"context"
	"net"

	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/interfaces"
)

type SessionService struct {
	locator      interfaces.GeoIPLocator
	sessionRepo  interfaces.SessionRepo
	sessionCache interfaces.SessionCache
}

func NewSessionService(locator interfaces.GeoIPLocator, repo interfaces.SessionRepo, cache interfaces.SessionCache) *SessionService {
	return &SessionService{
		locator:      locator,
		sessionRepo:  repo,
		sessionCache: cache,
	}
}

func (svc *SessionService) EnrichLocation(s *session.Session) {
	ip := net.ParseIP(s.DeviceInfo().IPAddr())

	location, _ := svc.locator.GetLocation(ip)
	s.UpdateLocation(location)
}

func (svc *SessionService) DeactivateSession(ctx context.Context, s *session.Session) error {
	s.Deactivate()

	err := svc.sessionCache.Del(ctx, s.ID())

	_, err = svc.sessionRepo.Update(ctx, s)
	if err != nil {
		return err
	}

	return nil
}
