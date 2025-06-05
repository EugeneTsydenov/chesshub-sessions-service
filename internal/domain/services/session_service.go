package services

import (
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/interfaces"
	"net"
)

type SessionService struct {
	locator interfaces.GeoIPLocator
}

func NewSessionService(locator interfaces.GeoIPLocator) *SessionService {
	return &SessionService{
		locator: locator,
	}
}

func (svc *SessionService) EnrichLocation(s *session.Session) {
	ip := net.ParseIP(s.DeviceInfo().IPAddr())

	location, _ := svc.locator.GetLocation(ip)
	s.UpdateLocation(location)
}
