package common_errors

import "fmt"

type UnknownIdentityError struct {
	CausedBy string
	Email    string
}

func (e UnknownIdentityError) Error() string {
	return fmt.Sprintf("Unknown identity %s caused by: %s", e.Email, e.CausedBy)
}
