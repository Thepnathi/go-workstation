package main

import "fmt"

// We can declare both parameter as a type

func multiply(x int, y int) int {
	return x * y
}

func multiplyTwo(x, y int) int {
	return x * y
}

func movieInformation(movie string) (int, bool, string) {
	var productionCost int
	var englishLanguage bool = false
	var director string = "Unknown"

	if movie == "Harry Potter" {
		productionCost, englishLanguage, director = 1870000, true, "James Bolt"
		return productionCost, englishLanguage, director
	}
	return productionCost, englishLanguage, director
}

func location(city string) (region, continent string) {
	switch city {
	case "Los Angeles", "LA", "Santa Monica":
		region, continent = "California", "North America"
	case "New York", "NYC":
		region, continent = "New York", "North America"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return
}

func funcFunc() {
	region, continent := location("Santa Monica")
	fmt.Printf("Matt lives in %s, %s", region, continent)
}
