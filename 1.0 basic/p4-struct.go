// Struct is a type container which contains properties
package main

type car struct {
	model      string
	name       string
	manufacter string
	price      int
}

func newStruct() {
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
}
