package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Init books var as a slice Book struct
var books []Book

// getBook : respond/request
func getBooks(w http.ResponseWriter, r *http.Request) {
	// set header value of content type to json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// getBook : get single book data
func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Get params
	// Loop through books with correct id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// Return a single book
	json.NewEncoder(w).Encode(&Book{})
}

// createBook : create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000)) // Mock ID - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// updateBook : update a book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// deleteBook : delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}
