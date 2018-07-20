package main

import (
	"fmt"
)

func slice() {
	mySlice := []int{0, 1, 2, 3, 4, 5}
	mySlice2 := mySlice[0:3]
	mySlice3 := mySlice2[1:]

	fmt.Println("The length of the slice is", len(mySlice), ". Last value is", mySlice[5])
	fmt.Println(mySlice, mySlice2, mySlice3)
	// The length of the slice is 6 . Last value is 5
	// [0 1 2 3 4 5] [0 1 2] [1 2]
}
