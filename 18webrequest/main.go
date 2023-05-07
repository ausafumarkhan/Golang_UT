package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to Web Request in Golang")
	response, err := http.Get(url)
	checkNilErr(err)
	fmt.Printf("Request is of type %T\n", response)
	defer response.Body.Close()

	dataByte, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
	content := string(dataByte)
	fmt.Println("The content of a web request is :", content)
}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
