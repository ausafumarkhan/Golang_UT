package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghh56n"

func main() {
	fmt.Println("Welcome to handling URLs in Golang")
	fmt.Println(myurl)

	// parsing --- > Extracting some information
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	//fmt.Println(result.RawPath)

	qparams := result.Query() //qparams -----> Query Parameters
	fmt.Printf("Type of query parameters is %T\n", qparams)
	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println("Param is ", val)
	}
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	anotherURL := partsOfUrl.String() // Reassembles the url into a valid URL
	fmt.Println(anotherURL)
}
