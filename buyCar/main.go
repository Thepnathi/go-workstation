package main

import (
	"fmt"
	"strconv"
)

type car struct {
	name  string
	price float64
}

type person struct {
	name   string
	budget float64
}

// Calculate if a person can afford a car. If not how much does he needs to save to be able to afford them

func affordThatCar(p person, carOne, carTwo car) (bool, person, car, float64) {
	budget := p.budget
	// we also create a struct for this
	var afford bool
	var per person
	var c car
	var diff float64

	switch {
	case budget >= carOne.price:
		afford = true
		per = p
		c = carOne
		diff = budget - carOne.price
	default:
		if carOne.price < carTwo.price {
			return false, p, carOne, carOne.price - budget
		} else {
			return false, p, carTwo, carTwo.price - budget
		}
	}
	return afford, per, c, diff
}

func main() {
	Eva := person{"Eva", 78000.78}

	Fiesta := car{"Ford Fiesta", 15500.00}
	F430 := car{"Farrari F430", 145999.99}

	afford1, owner1, car1, diff1 := affordThatCar(Eva, Fiesta, F430)

	remaining := strconv.FormatFloat(diff1, 'f', 2, 64)

	if afford1 == true {
		fmt.Printf("%s can afford %s with a remaining budget of £%s\n", owner1.name, car1.name, remaining)
	}
}
