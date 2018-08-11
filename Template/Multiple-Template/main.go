package main

import (
	"html/template"
	"os"
)

type Book struct {
	Name      string
	Author    string
	Published bool // check if a Book has been published yet
}

func main() {
	// henryPotter := Book{Name: "Henry Potterhead", Author: "James Smokes"} // published - false
	// scaryAlien := Book{Name: "Scary Alien AI", Author: "Cameroon Dickson", Published: true}

	tmpl, err := template.New("Using multiple templates").Parse(`
	{{define BookResource"}}
		<p>{{.Name}} by {{.Author}} {{if .Published}} (Published) {{else}}(Still being written){{end}}<p>
	{{end}}

	{{ define BookLoop"}}
		{{range .}}
			{{template BookResource" .}}
		{{else}}
			<p>No published Book yet</p>
		{{end}}
	{{end}}

	{{tempalte BookLoop" .}}
	`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, []Book{})
	if err != nil {
		panic(err)
	}

}
