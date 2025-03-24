package err

import (
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
