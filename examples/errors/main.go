package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func find(param string) error {
	if param == "equal" {
		return fmt.Errorf("buscar usuário: %w", ErrNotFound)
	}
	return fmt.Errorf("buscar usuário: %v", ErrNotFound)
}

func main() {
	// the error from services comes up
	err := find("equal")
	fmt.Println(err)

	err2 := find("not equal")
	fmt.Println(err2)

	// verify if err is the same that ErrNotFound
	fmt.Println("err is the type of ErrNotFound?", errors.Is(err, ErrNotFound))
	fmt.Println("err2 is the type of ErrNotFound?", errors.Is(err2, ErrNotFound))
}
