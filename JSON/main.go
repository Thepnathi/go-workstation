package main

import (
	"encoding/json"
	"fmt"
)

// Marshaling structs
type Article struct {
	Name       string
	AuthorName string
	draft      bool
}

func main() {
	article := Article{
		Name:       "JSON in Go",
		AuthorName: "Maria Knight",
		draft:      true,
	}
	data, err := json.Marshal(article)
	if err != nil {
		fmt.Println("Couldn't marshal article:", err)
	} else {
		// converts the byte data to string format
		fmt.Println(string(data))
	}

	data2, _ := json.MarshalIndent(article, "", " ")
	fmt.Println(string(data2))
}
