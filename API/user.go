package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Define the database
var db *gorm.DB
var err error

// User : The user model for the db
type User struct {
	// Let gorm know its the model for the DB
	gorm.Model
	Name  string
	Email string
}

// InitialMigration : Creates the table
func InitialMigration() {
	// connect to db
	// gorm will automatically create a table if it does not exist
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})
}

// AllUsers : Allows routes to access
// Returns all user in the SQLlite database
// r is the request object
func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()

	// Encode to JSON and return this as the response
	var users []User
	db.Find(&users)
	fmt.Println("{}", users)
	json.NewEncoder(w).Encode(users)
}

// NewUser : Create a new User for the db
func NewUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()
	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "New User Successfully Created")
}

// DeleteUser : Delete existing User from db
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	InitialMigration()
	fmt.Fprintf(w, "Delete User Endpoint Hit)")
}

// UpdateUser : Update the User
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	InitialMigration()
	fmt.Fprintf(w, "Update User Endpoint Hit")
}
