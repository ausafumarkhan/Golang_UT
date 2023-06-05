package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to maths in Golang")
	var a int = 3
	var b float64 = 4
	fmt.Println("The sum of both number will be", a+int(b))

	// Generating a Randiom number using a Crytopgraphy
	//rand.Seed(time.Now().UnixNano())
	// rand.New(rand.NewSource(99))
	// fmt.Println(rand.Intn(10))

	// Random from Crypto

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
