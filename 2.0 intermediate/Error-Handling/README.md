# Error Handling in Go

Error handling is a convention in Golang.

Error handling is extremely important in our world since we 
relied on external resources, such as network connection and disk
availability.

Go error is an INTERFACE, it is any type that implementds the 
Error method 

type error interface {
    Error() string
}

https://blog.golang.org/error-handling-and-go
