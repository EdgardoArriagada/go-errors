package main

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

func run() error {
	return &Temporary{
		Temporary: true,
		message:   "didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
