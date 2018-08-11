package main

import (
	"html/template"
	"os"
)

type product struct {
	Price    float64
	Quantity float64
}

func multiply(a, b float64) float64 {
	return a * b
}

func main() {
	// register the func for with our template

	tmpl, _ := template.New("Example-Pipe").Parse(
		"Price: ${{printf \"%.2f\" .}}\n",
	)
	tmpl.Execute(os.Stdout, 22.3) // Price: $22.30

	tmpl2, _ := template.New("Chain pipeline").Parse(
		"Price: ${{. | printf \"%.2f\"}}\n",
	)
	tmpl2.Execute(os.Stdout, 12.8)

	// Two pipe example with func

	tmplX := template.New("Func Pipeline")
	tmplX.Funcs(template.FuncMap{"multiply": multiply})

	// After one pipe we use | to connect with another pipe
	tmplX, err := tmplX.Parse(
		"Price: ${{ multiply .Price .Quantity | printf \"%.2f\"}}\n",
	)
	if err != nil {
		panic(err)
	}

	err = tmplX.Execute(os.Stdout, product{
		Price:    5.1,
		Quantity: 10,
	}) // Price: $51.00
	if err != nil {
		panic(err)
	}

	// Pipe with template variable

	tmpl3, _ := template.New("Chain pipeline").Parse(`
		{{ $total := multiply .Price .Quantity}}
		"Price: ${{ printf "%.2f" $total}}\n",
	`)
	tmpl3.Execute(os.Stdout, 12.8)
}
