package main

import (
	"errors"
	"example/user/go-errors/cmd/err"
	"fmt"
)

func main() {
	normalError := errors.New("normal error")
	temporaryError := err.NewTemporaryError("some temporary error")

	fmt.Println("with normal error", err.IsTemporary(normalError))
	fmt.Println("with temporary error", err.IsTemporary(temporaryError))

	// printing messages
	fmt.Println("with normal error:", normalError)
	fmt.Println("with temporary error:", temporaryError)
}
