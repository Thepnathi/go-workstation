package main

import (
	"fmt"
)

// EXAMPLE 1
// This is not like traditional OOP because the area method should be inside the Rectangle struct
// In case of OOP it would be inside the a rectangle class
// They are both two independent thing

// Rectangle information
type Rectangle struct {
	width, height float64
}

func area(r Rectangle) float64 {
	return r.height * r.width
}

// EXAMPLE 2
// A method has a type

func main() {
	r1 := Rectangle{10, 5}
	fmt.Println(area(r1))
}
