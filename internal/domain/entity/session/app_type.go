package session

import "fmt"

type AppType int

const (
	AppTypeChesshubWeb AppType = iota
	AppTypeChesshubMobile
	AppTypeChesshubDesktop
	AppTypeChesshubTablet
)

var (
	AppTypeName = map[AppType]string{
		AppTypeChesshubWeb:     "ChesshubWeb",
		AppTypeChesshubMobile:  "ChesshubMobile",
		AppTypeChesshubDesktop: "ChesshubDesktop",
		AppTypeChesshubTablet:  "ChesshubTablet",
	}
	AppTypeValue = map[string]AppType{
		"ChesshubWeb":     AppTypeChesshubWeb,
		"ChesshubMobile":  AppTypeChesshubMobile,
		"ChesshubDesktop": AppTypeChesshubDesktop,
		"ChesshubTablet":  AppTypeChesshubTablet,
	}
)

func (a *AppType) String() string {
	return AppTypeName[*a]
}

func (a *AppType) IsValid() bool {
	switch *a {
	case AppTypeChesshubWeb, AppTypeChesshubMobile, AppTypeChesshubDesktop, AppTypeChesshubTablet:
		return true
	default:
		return false
	}
}

func (a *AppType) Scan(value any) error {
	i, ok := value.(int64)
	if !ok {
		return fmt.Errorf("cannot scan app type: %v", value)
	}

	t := AppType(i)
	if !t.IsValid() {
		return fmt.Errorf("invalid app type: %d", i)
	}

	return nil
}
