package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello")

	// ======================
	// 1. Variable declaration

	const WEBSITE string = "www.fsc.com"
	const superSmall float64 = 0.012345
	var number int = 123
	var measure float64 = 1.23323
	var name string = "Thepnathi"
	var notTrue bool = false
	shortDeclaOperator := "Go detects this type automatically"

	// Will through a space
	fmt.Println(number, measure)
	fmt.Println("Hello, " + name + "\n" + shortDeclaOperator + ". Check this site " + WEBSITE)
	fmt.Println(len(name))
	fmt.Println(notTrue)
	// Define how many dp to print
	fmt.Printf("%.3f\n", superSmall)
	// Check data-type
	fmt.Printf("%T \n", superSmall)

	// ======================
	// 2. La Loop

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for k := 0; k < 3; k++ {
		fmt.Println(k)
	}
	fmt.Println("")

	infinite := true
	for s := 0; infinite; s++ {
		infinite = false
	}

	// ======================
	// 3. Control Flow

	currentBalance := 3456.03

	if currentBalance == 0 {
		fmt.Println("You have no money!")
	} else if currentBalance < 1000 {
		fmt.Println("You have low balance!")
	} else {
		fmt.Println("Keep saving!")
	}

	month := 01

	switch month {
	case 01:
		fmt.Println("January")
	case 02:
		fmt.Println("February")
	default:
		fmt.Println("No idea")
	}

	// ======================
	// 4. Data Structure

	// This is an array
	var colours [2]string
	colours[0] = "Yellow"
	colours[1] = "Green"
	countries := [3]string{"Albania", "Italy", "UK"}

	for i, country := range countries {
		fmt.Println(country, i)
	}

	// Not print index
	for _, element := range countries {
		fmt.Println(element)
	}

	fmt.Println("")
	// This is a slice
	fruits := []string{"Orange", "Mango", "Bananan"}
	fruits = append(fruits, "Pear") // Will return a new slice

	productStock := make(map[string]int)

	productStock["Orange"] = 102
	productStock["Banana"] = 123
	productStock["Strawberry"] = 0
	productStock["Mango"] = 78
	productStock["Durian"] = 122

	delete(productStock, "Durian")

	fmt.Println(productStock)
	fmt.Println("Orange quantity: " + strconv.Itoa(productStock["Orange"]) + "\n")

	for fruit, quantity := range productStock {
		fmt.Println("fruit: ", fruit, "quantity: ", quantity)
	}
	fmt.Println("")

	// ======================
	// 4. Functions (see function.go)

	sqr := areaOfSquare(2, 10)
	fmt.Println(sqr)

	result, err := squareRoot(16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	resultTwo, err := squareRoot(-1111)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultTwo)
	}
	fmt.Println("")

	// ======================
	// 4. Type (Similar to OOP object) see things.go

	buddy := cat{name: "Buddy", age: 7, eyeColor: "Brown", meow: true}
	fmt.Println(buddy)
	fmt.Println(buddy.name + "\n")

	// ======================
	// 4. Pointers

	// This will get incremented and discarded
	l := 89
	fmt.Println(&l) // Give us a pointer to aj
	errorIncrement(l)
	fmt.Println(l)

	// We have to pass on the memory address and alternate it
	m := 99
	properIncrement(&m)
	fmt.Println(m)
}
