package main

import (
	"errors"
	"example/user/go-errors/cmd/nerr"
	"fmt"
)

func function2() error {
	return nerr.NewBadRequestError("bad request error")
}

func function1() error {
	if err := function2(); err != nil {
		return errors.Join(err, errors.New("function2 failed"))
	}

	return nil
}

func function0() error {
	if err := function1(); err != nil {
		return errors.Join(err, errors.New("function1 failed"))
	}

	return nil
}

func main() {
	normalError := errors.New("normal error")
	temporaryError := nerr.NewTemporaryError("some temporary error")

	fmt.Println("with normal error", nerr.IsTemporary(normalError))
	fmt.Println("with temporary error", nerr.IsTemporary(temporaryError))

	// printing messages
	fmt.Println("with normal error:", normalError)
	fmt.Println("with temporary error:", temporaryError)

	badRequestError := function0()
	fmt.Println("with bad request error:", badRequestError)
}
