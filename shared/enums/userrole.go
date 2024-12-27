package enums

import "errors"

type Role int

const (
	EndUser Role = 1
	Admin   Role = 2
)

func (r Role) GetRoleName() (string, error) {
	switch r {
	case EndUser:
		return "EndUser", nil
	case Admin:
		return "Admin", nil
	default:
		return "", errors.New("no role type found")
	}

}
