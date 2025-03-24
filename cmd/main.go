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

func IsTemporary(err error) bool {
	te, ok := err.(*Temporary)
	return ok && te.Temporary
}

func main() {
	normalError := errors.New("normal error")
	temporaryError := &Temporary{Temporary: true, message: "temporary error"}

	fmt.Println("with normal error", IsTemporary(normalError))
	fmt.Println("with temporary error", IsTemporary(temporaryError))
}
