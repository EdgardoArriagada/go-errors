package nerr

import (
	"fmt"
)

type Error struct {
	Temporary  bool
	BadRequest bool
	message    string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s", e.message)
}

func NewTemporaryError(message string) *Error {
	return &Error{Temporary: true, message: message}
}

func NewBadRequestError(message string) *Error {
	return &Error{BadRequest: true, message: message}
}

func IsTemporary(err error) bool {
	te, ok := err.(*Error)
	return ok && te.Temporary
}

func IsBadRequest(err error) bool {
	te, ok := err.(*Error)
	return ok && te.BadRequest
}
