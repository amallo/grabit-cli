package core_errors

import "fmt"

const (
	ErrNotFound = "NOT_FOUND"
)

type Error interface {
	Code() string
	CausedBy() error
	Error() string
}

type commonError struct {
	code     string
	causedBy error
}

func Err(code string, causedBy error) Error {
	return commonError{code: code, causedBy: causedBy}
}

func (e commonError) Code() string {
	return e.code
}

func (e commonError) CausedBy() error {
	return e.causedBy
}

func (e commonError) Error() string {
	return fmt.Sprintf("%s: caused by %s", e.code, e.causedBy)
}

func NotFoundErr(causedBy error) Error {
	return commonError{code: ErrNotFound, causedBy: causedBy}
}
