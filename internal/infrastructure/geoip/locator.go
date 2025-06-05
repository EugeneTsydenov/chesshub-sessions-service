package geoip

import (
	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
	"net"
)

type GeoIPLocator struct {
	database *Database
}

func NewGeoIPLocator(db *Database) *GeoIPLocator {
	return &GeoIPLocator{
		database: db,
	}
}

func (l *GeoIPLocator) GetLocation(ip net.IP) (*session.Location, error) {
	record, err := l.database.Conn().Get_all(ip.String())
	if err != nil {
		return nil, err
	}

	city := record.City
	country := record.Country_long

	location := session.NewLocation(nilIfEmpty(city), nilIfEmpty(country))

	return location, nil
}

func nilIfEmpty(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
