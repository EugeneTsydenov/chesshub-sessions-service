package geoip

import (
	"context"
	"github.com/ip2location/ip2location-go/v9"
)

type Database struct {
	conn *ip2location.DB
}

func New(path string) (*Database, error) {
	db, err := ip2location.OpenDB(path)
	if err != nil {
		return nil, err
	}
	return &Database{
		conn: db,
	}, nil
}

func (d *Database) Conn() *ip2location.DB {
	return d.conn
}

func (d *Database) Shutdown(_ context.Context) error {
	d.conn.Close()
	return nil
}
