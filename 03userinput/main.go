package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for pizza")
	// comma ok || err err

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for your, ", input)
	fmt.Printf("Type of rating is %T\n", input)
}
