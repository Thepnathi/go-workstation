package main

import "fmt"

func changeFruit(fruit string) {
	fruit = "pear"
}

func pointerFruit(fruit *string) {
	*fruit = "pear"
}

func main() {
	apple := "apple"
	changeFruit(apple)
	fmt.Println(apple) // apple

	strawberry := "strawberry"
	pointerFruit(&strawberry) // Needs &
	fmt.Println(strawberry)   // pear
}
