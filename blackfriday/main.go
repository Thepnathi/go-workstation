package main

import (
	"fmt"

	"github.com/russross/blackfriday"
)

func main() {
	markdown := []byte(`
# This is a header
* This 
* is 
* a 
* list
	`)
	html := blackfriday.MarkdownBasic(markdown)
	fmt.Println(string(html))
}
