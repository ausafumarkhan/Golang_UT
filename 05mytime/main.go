package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time handling code")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2022, time.February, 14, 13, 45, 50, 0, time.UTC)
	fmt.Println(createdDate)
}
