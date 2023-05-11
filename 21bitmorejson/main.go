package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` // json value will reflected in output, e.g. Name will be reflected as coursename
	Price    int
	Platform string   `json:"website"`        // Platfrom will be reflected as website in Json data
	Password string   `json:"-"`              // "-" is used to hide that field in the API or encoded Json data
	Tags     []string `json:"tags,omitempty"` // omitempty will remove all the spaces
}

func main() {
	fmt.Println("Welcome to JSON video")
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"React Bootcamp", 299, "Youtube", "abc123", []string{"beginner", "web-dev"}},
		{"Python Bootcamp", 199, "Udemy", "xyz123", []string{"beginner", "scripting"}},
		{"Golang Bootcamp", 399, "LcodeOnline", "hit123", []string{"beginner", "backend-dev"}},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", finalJson)
}
