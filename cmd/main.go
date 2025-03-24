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

func function22() error {
	// some external error
	return errors.New("some external error")
}

func function11() error {
	// wrap it making it temporary
	if err := function22(); err != nil {
		error := errors.New("at function 11")
		return errors.Join(err, error, nerr.TemporaryError)
	}

	return nil
}

func function00() error {
	if err := function11(); err != nil {
		return errors.Join(err, errors.New("at function00"))
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

	approach2Error := function00()
	fmt.Println("le approach2Error", approach2Error)

	if errors.Is(approach2Error, nerr.TemporaryError) {
		fmt.Println("is temporary error!")
	}
}
