package main

// Interface to access data from a file and a database

type Product struct {
	name        string
	description string
}

type ProductCatalogue interface {
	All() ([]Product, error)
	Find(string) (Product, error)
}

type FileProductCatalogue struct {
	// ... e.g. file location
}

func (c FileProductCatalogue) All() ([]Product, error) {
	//
}

type DBProductCatalogue struct {
	// e.g. database connection
}

func main() {
	var myCatalogue ProductCatalogue

}
