package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs in Golang")
	// No inheritance in Golang, No Super or parent
	ausaf := User{"ausaf", "ausaf@go.com", true, 26}
	fmt.Println(ausaf)
	fmt.Printf("Ausaf details are %+v\n", ausaf)
	fmt.Printf("Name is %v and Email is %v\n", ausaf.Name, ausaf.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
