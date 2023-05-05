package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Golang")
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}
	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}
	// We can replace index with '_' if we dont want to use it

	rogueValue := 1

	for rogueValue < 10 {
		if rogueValue == 2 {
			goto lco
		}
		if rogueValue == 5 {
			rogueValue++
			continue
		}
		fmt.Println("Value is", rogueValue)
		rogueValue++
	}
lco:
	fmt.Println("Jumping at learnCodeOnline")
}
