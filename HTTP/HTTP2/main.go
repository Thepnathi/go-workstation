package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.ServeMux{}
	mux.HandleFunc("/articles/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from /articles/")
	})
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from /users")
	})
	http.ListenAndServe(":3000", mux)
}
