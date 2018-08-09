package main

import "fmt"

type User struct {
	Name string
}

func (u User) isAdmin() bool { return false }
func (u User) displayName() string {
	return u.Name
}

// Admin Embeded type is similar to inheritanc in OOP
type Admin struct {
	User
}

func (a Admin) isAdmin() bool { return true }
func (a Admin) displayName() string {
	return "An admin has appeared! " + a.User.displayName()
}
func main() {
	u := User{"Normal User"}
	fmt.Println(u.Name)          // Normal User
	fmt.Println(u.displayName()) // Normal User
	fmt.Println(u.isAdmin())     // false

	a := Admin{User{"Robot"}}
	fmt.Println(a.Name)          // Robot
	fmt.Println(a.User.Name)     // Robot
	fmt.Println(a.displayName()) // An admin has appeared! Robot
	fmt.Println(a.isAdmin())     // false
}
