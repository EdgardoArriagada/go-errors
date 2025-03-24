package main

import (
	"errors"
	"fmt"
)

type Temporary struct {
	Temporary bool
	message   string
}

func (e *Temporary) Error() string {
	return fmt.Sprintf("%s", e.message)
}

func NewTemporaryError(message string) *Temporary {
	return &Temporary{Temporary: true, message: message}
}

func IsTemporary(err error) bool {
	te, ok := err.(*Temporary)
	return ok && te.Temporary
}

func main() {
	normalError := errors.New("normal error")
	temporaryError := NewTemporaryError("some temporary error")

	fmt.Println("with normal error", IsTemporary(normalError))
	fmt.Println("with temporary error", IsTemporary(temporaryError))
}
