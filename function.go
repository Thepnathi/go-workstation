package main

import (
	"errors"
	"math"
)

func areaOfSquare(width int, height int) int {
	return width * height
}

// Go can return multiple types
// Go does not have exception handling
func squareRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("Undefined negative numbers")
	}
	return math.Sqrt(num), nil
	// nil to prevent returning error
}

func errorIncrement(x int) {
	x++
}

// Excepts pointer using *
func properIncrement(x *int) {
	// De-reference the address
	// We do not want to increment the memory address
	*x++
}
