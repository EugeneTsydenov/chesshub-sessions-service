package interfaces

import (
	"net"

	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
)

type GeoIPLocator interface {
	GetLocation(ip net.IP) (*session.Location, error)
}
