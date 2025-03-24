package main

import (
	"errors"
	"example/user/go-errors/cmd/nerr"
	"fmt"
)

func function2() error {
	// some external error
	return errors.New("some external error")
}

func function1() error {
	// wrap it making it temporary
	if err := function2(); err != nil {
		error := errors.New("at function 1")
		return errors.Join(err, error, nerr.TemporaryError)
	}

	return nil
}

func function0() error {
	if err := function1(); err != nil {
		return errors.Join(err, errors.New("at function 0"))
	}

	return nil
}

func main() {
	normalError := errors.New("normal error")
	temporalError := function0()
	// log each error
	fmt.Println("le normal error:", normalError)
	fmt.Println("le temporal error:", temporalError)

	if errors.Is(temporalError, nerr.TemporaryError) {
		fmt.Println("'temporalError' is temporary error!")
	} else {
		fmt.Println("'temporalError' is NOT temporary error.")
	}

	if errors.Is(normalError, nerr.TemporaryError) {
		fmt.Println("'normalError' is temporary error!")
	} else {
		fmt.Println("'normalError' is NOT temporary error.")
	}
}
