package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Init the Router
	router := mux.NewRouter()

	// Mock data - @todo - implement DB
	books = append(books, Book{ID: "1", Isbn: "8282", Title: "Book one", Author: &Author{Firstname: "John", Surname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "1233", Title: "Book Two", Author: &Author{Firstname: "Max", Surname: "Jso"}})
	books = append(books, Book{ID: "3", Isbn: "8229", Title: "Book Three", Author: &Author{Firstname: "Eva", Surname: "Knight"}})

	// Route Handler / Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	// log.Fatal will throw err if it fails
	log.Fatal(http.ListenAndServe(":9000", router))
}
