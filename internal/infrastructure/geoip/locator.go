package geoip

import (
	"log"
	"net"
	"strings"

	"github.com/EugeneTsydenov/chesshub-sessions-service/internal/domain/entity/session"
)

type Locator struct {
	database *Database
}

func NewLocator(db *Database) *Locator {
	return &Locator{
		database: db,
	}
}

func (l *Locator) GetLocation(ip net.IP) (*session.Location, error) {
	record, err := l.database.Conn().Get_all(ip.String())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	city := record.City
	country := record.Country_long
	location := session.NewLocation(nilIfInvalid(city), nilIfInvalid(country))

	return location, nil
}

func isValidLocation(location string) bool {
	return !strings.Contains(location, "IP2Location")
}

func nilIfInvalid(s string) *string {
	if s == "" {
		return nil
	}

	if !isValidLocation(s) {
		return nil
	}

	return &s
}
