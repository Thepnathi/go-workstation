package main

import (
	"fmt"
)

type ErrInvalidStatusCode int

func (code ErrInvalidStatusCode) Error() string {
	return fmt.Sprintf("Expected 200 ok, but got %d", code)
}

func main() {
	statusCode := 404

	if statusCode != 200 {
		fmt.Println(ErrInvalidStatusCode(statusCode))
	}
}
