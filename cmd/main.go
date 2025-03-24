package main

import (
	"errors"
	"example/user/go-errors/cmd/nerr"
	"fmt"
)

func externalFunction() error {
	return errors.New("some external error")
}

func function2() error {
	errorCtx := errors.New("at function2")
	// wrap it making it temporary
	if err := externalFunction(); err != nil {

		return errors.Join(nerr.TemporaryError, err, errorCtx,
			errors.New("failed calling externalFunction"),
		)
	}

	return nil
}

func function1() error {
	errorCtx := errors.New("at function1")
	if err := function2(); err != nil {
		return errors.Join(err, errorCtx, errors.New("calling function1"))
	}

	return nil
}

func main() {
	normalError := errors.New("normal error")
	temporalError := function1()
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
