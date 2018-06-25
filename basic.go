package main

import "fmt"

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

}
