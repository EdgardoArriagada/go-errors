package main

import (
	"errors"
	"example/user/go-errors/cmd/nerr"
	"fmt"
)

func main() {
	normalError := errors.New("normal error")
	temporaryError := nerr.NewTemporaryError("some temporary error")

	fmt.Println("with normal error", nerr.IsTemporary(normalError))
	fmt.Println("with temporary error", nerr.IsTemporary(temporaryError))

	// printing messages
	fmt.Println("with normal error:", normalError)
	fmt.Println("with temporary error:", temporaryError)
}
