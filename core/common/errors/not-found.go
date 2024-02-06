package common_errors

import "fmt"

type NotFoundError struct {
	CausedBy string
	Category string
	Id       string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s not found", e.Category)
}
