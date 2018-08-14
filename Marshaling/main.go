package main

import (
	"encoding/json"
	"fmt"
)

// Custom JSON key, we can change what JSON

type Article struct {
	Name string
}

type ArticleCollection struct {
	Articles []Article `json:"articles"`
	Total    int       `json:"total"`
}

func main() {
	p1 := Article{Name: "JSON in Go"}
	p2 := Article{Name: "Marshaling is easy"}
	articles := []Article{p1, p2}

	collection := ArticleCollection{
		Articles: articles,
		Total:    len(articles),
	}
	data, err := json.MarshalIndent(collection, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
