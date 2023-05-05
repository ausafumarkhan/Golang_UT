package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointers in golang")
	var ptr *int
	fmt.Println("The default value of a pointer is:", ptr)

	myNumber := 23
	var newPtr = &myNumber
	fmt.Println("The value of a pointer is", *newPtr)
	fmt.Println("New value of a pointer is", *newPtr+2)

}
