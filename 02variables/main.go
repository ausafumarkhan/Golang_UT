package main

import "fmt"

const LoginToken string = "ghnabvroshnbdbsjua"

func main() {
	var username string = "ausaf"
	fmt.Println(username)
	fmt.Printf("Variable is of type %T \n", username)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type %T \n", smallVal)

	var smallFloat float32 = 255.6445673
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type %T \n", smallFloat)

	// Default value and aliases
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type %T \n", anotherVar)

	// Implicit type
	var website = "hisashitech.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type %T \n", website)

	// no var
	numberOfUser := 30000
	fmt.Println(numberOfUser)

	fmt.Printf("Variable is of %T \n", LoginToken)
}
