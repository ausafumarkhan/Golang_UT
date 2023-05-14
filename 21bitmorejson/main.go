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
	//EncodeJson()
	DecodeJson()
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

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "React Bootcamp",
		"Price": 299,
		"website": "Youtube",
		"tags": ["beginner","web-dev"]
	}
	`)
	var lcocourse course

	// We can validate the json data (format) using inbuilt method

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json was Valid")
		json.Unmarshal(jsonDataFromWeb, &lcocourse)
		fmt.Printf("%#v\n", lcocourse)
	} else {
		fmt.Println("Json was not valid")
	}
	// Some cases we just want to add data into key value pair
	// key can be a string but there is no guarantee of value's datatype so we use "interface"

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v. Type of value is %T\n", k, v, v)
	}

}
