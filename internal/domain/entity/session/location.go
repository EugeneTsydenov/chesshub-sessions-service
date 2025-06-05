package session

type Location struct {
	city    *string
	country *string
}

func NewLocation(city, country *string) *Location {
	return &Location{
		city:    city,
		country: country,
	}
}

func (s *Location) City() *string {
	return s.city
}

func (s *Location) Country() *string {
	return s.country
}
