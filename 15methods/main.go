package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs in Golang")
	// No inheritance in Golang, No Super or parent
	ausaf := User{"ausaf", "ausaf@go.com", true, 31}
	faraaz := User{"faraaz", "faraaz@go.com", false, 25}
	fmt.Println(ausaf)
	fmt.Printf("Ausaf details are %+v\n", ausaf)
	fmt.Printf("Name is %v and Email is %v\n", ausaf.Name, ausaf.Email)

	ausaf.GetStatus()
	faraaz.GetStatus()

	faraaz.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}

// We need to modify the above the code for updating the value of each particular user
