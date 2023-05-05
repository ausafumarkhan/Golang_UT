package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in Golang")
	fruitList := []string{"Apple", "Banana", "Grapes"} // var fruitList = []string{"Apple", "Banana", "Grapes"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)
	fmt.Println(fruitList)

	fruitList = append(fruitList, "Peach", "Mango")
	fmt.Println(fruitList)

	// Creating slices
	fmt.Println(fruitList[1:])
	fmt.Println(fruitList[1:3])
	fmt.Println(fruitList[:3])

	//HighScores
	var highScores = make([]int, 4) // highScores := make([]int, 4)
	highScores[0] = 435
	highScores[1] = 234
	highScores[2] = 782
	highScores[3] = 729
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)

	// How to remove a value from slice based on index in Golang
	var courses = []string{"reactjs", "ruby", "javascript", "python", "golang"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
