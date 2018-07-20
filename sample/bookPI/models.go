package main

// Book structs (Model)
// Class use for OOP :)
// *Author has its own struct
type Book struct {
	ID     string  `json: "id"`
	Isbn   string  `json: "isbn"`
	Title  string  `json: "title"`
	Author *Author `json: "author"`
}

// Author struct
type Author struct {
	Firstname string `json: "firstname"`
	Surname   string `json: "lastname"`
}
