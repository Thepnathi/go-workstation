package main

import (
	"fmt"
	"html/template"
	"os"
)

type book struct {
	Name      string
	Author    string
	Published bool // check if a book has been published yet
}

func (b book) Byline() string {
	return fmt.Sprintf("Written by %s", b.Author)
}

func main() {
	henryPotter := book{Name: "Henry Potterhead", Author: "James Smokes"}

	tmpl, err := template.New("TemplateName").Parse("<h1>Hello {{.}}</h1>\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, "World")
	if err != nil {
		panic(err)
	}
	// example 2 - similar technique can be applied to maps
	fmt.Println(henryPotter.Published) // false
	tmpl2, err := template.New("Accessing a struct template").Parse("{{.Name}} by {{.Author}} {{if .Published}} (Published) {{else}}(Still being written){{end}}\n")
	if err != nil {
		panic(err)
	}

	err = tmpl2.Execute(os.Stdout, henryPotter)
	if err != nil {
		panic(err)
	}

	// example 3 - using niladic function (see Go docs)
	tmpl3, err := template.New("Using niladic function").Parse("{{.Byline}}\n")
	if err != nil {
		panic(err)
	}
	err = tmpl3.Execute(os.Stdout, henryPotter)
	if err != nil {
		panic(err)
	}

	// example 4 - Loops with range
	// this example returns the else option since it has to use slice/map as the data
	tmpl4, err := template.New("Using niladic function").Parse(`
	<h1>Loops commence!</h1>
	{{range .}}
		<p>{{.Name}} by {{.Author}}</p>
	{{else}}
		<p>No published book yet</p>
	{{end}}
	`)
	if err != nil {
		panic(err)
	}
	err = tmpl4.Execute(os.Stdout, []book{})
	if err != nil {
		panic(err)
	}
}
