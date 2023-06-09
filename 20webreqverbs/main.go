package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb in Golang")
	PerformGetResponse()
	//PerfromPostJsonRequest()
	//PerformPostFormRequest()
}
func PerformGetResponse() {
	const myurl string = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Response is", response)
	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content length is:", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)

	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is", byteCount)
	fmt.Println(responseString.String()) //It will give all the hold by "responseString" similar to "string(content)"
	//fmt.Println(content)
	//fmt.Println(string(content))

}
func PerfromPostJsonRequest() {
	const myurl = "http://localhost:8000/post"
	// fake json payload
	requestBody := strings.NewReader(`
		{
			"cousrsename":"Let's start with Golang",
			"price":0,
			"platform":"youtube"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// for postform

	data := url.Values{}
	data.Add("firstname", "ausaf")
	data.Add("lastname", "khan")
	data.Add("email", "ausaf@go.dev")
	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	content, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	fmt.Println(string(content))
}
