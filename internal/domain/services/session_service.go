package services

import (
	"context"
	"net"

	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/interfaces"
)

type SessionService struct {
	locator interfaces.GeoIPLocator
	repo    interfaces.SessionRepo
}

func NewSessionService(locator interfaces.GeoIPLocator, repo interfaces.SessionRepo) *SessionService {
	return &SessionService{
		locator: locator,
		repo:    repo,
	}
}

func (svc *SessionService) EnrichLocation(s *session.Session) {
	ip := net.ParseIP(s.DeviceInfo().IPAddr())

	location, _ := svc.locator.GetLocation(ip)
	s.UpdateLocation(location)
}

func (svc *SessionService) DeactivateSession(ctx context.Context, s *session.Session) error {
	s.Deactivate()

	_, err := svc.repo.Update(ctx, s)
	if err != nil {
		return err
	}

	return nil
}
