package interfaces

import (
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
	"net"
)

type GeoIPLocator interface {
	GetLocation(ip net.IP) (*session.Location, error)
}
