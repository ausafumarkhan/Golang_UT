package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to Switch and Case in Golang")
	rand.New(rand.NewSource(99))

	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of Dice number is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice number is 1 & you can open to play")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
	case 4:
		fmt.Println("You can move 4 spot")
	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("You can move 6 spot and play again")
	default:
		fmt.Println("What was that")
	}

}
