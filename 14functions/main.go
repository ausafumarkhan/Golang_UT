package main

import "fmt"

func main() {
	fmt.Println("Welcome to the functions on golang")
	greater()

	result := adder(2, 3)
	fmt.Println("The result of addition is:", result)

	proResult, proMessage := proAdder(2, 4, 5, 6, 7)
	fmt.Println("Pro Result is:", proResult)
	fmt.Println("Pro Message is:", proMessage)
}

func greater() {
	fmt.Println("Hello from Golang")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi proResult function"
}
