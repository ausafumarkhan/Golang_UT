package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Peach"
	fruitList[3] = "Grape"
	fmt.Println("The lsit of fruits is: ", fruitList)
	fmt.Printf("The list of fruits is %v\n", fruitList)

	var vegList = [3]string{"carrot", "potato", "onion"}
	fmt.Println("The list of vegetables:", vegList)
}
