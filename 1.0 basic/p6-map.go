package main

import "fmt"

func mapFunc() {
	// Data structure to assign key and value
	// Similar to JavaScript or Python's Dictionary

	starWarsYears := map[string]int{
		"A New Hope":              1977,
		"The Empire Strikes Back": 1980,
		"Return of the Jedi":      1983,
		"Attack of the Clone":     2002,
		"Revenge of the Sith":     2005,
	}

	// Add new value into map
	starWarsYears["The Force Awakens"] = 2015

	// Get a corresponding key
	fmt.Println(starWarsYears["A New Hope"])

	fmt.Println(len(starWarsYears))

	// REMEMBER - map will output in any order
	for title, year := range starWarsYears {
		fmt.Println(title, "was released in ", year)
	}

	// Delete from a map
	delete(starWarsYears, "The Force Awakens")
}
