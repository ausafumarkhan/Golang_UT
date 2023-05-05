package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Maps in golang")
	var languages = make(map[string]string)
	languages["js"] = "Javascript"
	languages["py"] = "Python"
	languages["rb"] = "Ruby"
	fmt.Println("List of all languages: ", languages)
	fmt.Println("js is short for :", languages["js"])

	// Loops are very interesting in Golang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
