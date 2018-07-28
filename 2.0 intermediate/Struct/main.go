package main

import "fmt"

// struct is a type which contains name field
// Field which contains capital letter are accessible outside of the package
// it will be public rather than private

// Animal - information about an  animal
type Animal struct {
	Name     string
	NumLeg   int
	Location []string
}

type car struct {
	model      string
	name       string
	manufacter string
	price      int
}

func main() {
	// Instance of animal

	lion := Animal{
		Name:   "Lion",
		NumLeg: 4,
	}

	lion.Location = []string{
		"India",
		"Madagascar",
		"Thailand",
	}

	// ways to define struct
	var GT car
	GT.model = "UI7889"
	GT.name = "Ford GT"
	GT.manufacter = "Ford"
	GT.price = 2500000

	Veneno := car{"UJ89", "Lamborghini Veneno Roaster", "Lamborghini", 23742882}

	// Anonymous struct

	tommy := struct {
		name string
		age  int
	}{"Tommy", 23}

	fmt.Println(lion)
	fmt.Println(Veneno)
	fmt.Println(tommy)
}
