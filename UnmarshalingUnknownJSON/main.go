package main

import (
	"encoding/json"
	"fmt"
)

// Unmarshal - For when we don't know exact JSON format
// When we selecting a type for the unknown input, we
// can use a blank interface type interface{}

func FooJSON(input string) {
	data := map[string]interface{}{}
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		panic(err)
	}

	foo, _ := data["foo"]

	switch foo.(type) {
	case float64:
		fmt.Printf("Float %f\n", foo)
	case string:
		fmt.Printf("String %s\n", foo)
	default:
		fmt.Printf("Somehting else\n")
	}
}

func main() {
	FooJSON(`{
		"foo": 123
		}`)
}
