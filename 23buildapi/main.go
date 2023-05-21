package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Models for courses - file
// - file means it will go to separate file
type Course struct {
	CourseId    string   `json:"courseid"`
	CourseName  string   `json:"coursename"`
	CoursePrize int      `json:"price"`
	Authors     *Authors `json:"author"` // Custom Type We use pointer to define its datatype
}

type Authors struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Creating a fake database using slice as we are not use any database
var courses []Course

/*
There are some helper methods which help to perform certain task as encrytion of password,
not allowing user to pull from database without unique courseID etc.
Called as middleware or helper method
*/

func (c *Course) IsEmpty() bool {
	//return c.CourseId == "" && c.CourseName == "" // We can create more such methods
	return c.CourseName == ""
}

func main() {
	fmt.Println("API-LearnCodeOnline.in")

	// Create a router using mux
	r := mux.NewRouter()

	// Seeding of data
	courses = append(courses, Course{CourseId: "102", CourseName: "Golang-Basic", CoursePrize: 299,
		Authors: &Authors{Fullname: "Ausaf Umar", Website: "basic@go.dev"}})
	courses = append(courses, Course{CourseId: "105", CourseName: "Golang-Advanced", CoursePrize: 499,
		Authors: &Authors{Fullname: "Ausaf Umar", Website: "advanced@go.dev"}})

	// Routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourse).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCOurse).Methods("GET")
	r.HandleFunc("/course", CreateOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", DeleteOneCourse).Methods("DELETE")

	// Listem to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Controller - file

/*
	"Serve Home route"

Always governed by 2 things--
1. w : writer -- writing a response for request and 2. r : reader -- getting request from sender
*/
func serveHome(w http.ResponseWriter, r *http.Request) {
	// Writing a response
	w.Write([]byte("<h1>Welcome to Build Api by LearnCodeOnline </h1>"))
}

// Creating aonther route to throw all the values from the databases

func getAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all the courses")
	// Method to set the headers, it is needed sometime to set explicit headers
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCOurse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")
	// Grab course ID from request
	params := mux.Vars(r)

	// loop trhough the courses, find id and return the response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given ID")
	return
}

func CreateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one Course")
	w.Header().Set("Content-Type", "application/json")
	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about: {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//TODO: check only if title is duplicate
	// loop, title matches with course.coursename, JSON

	// Genearte unique ID, string
	// append course into courses

	rand.New(rand.NewSource(100))

	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one Course")
	w.Header().Set("Content-Type", "application/json")

	// Grab ID from request

	params := mux.Vars(r)

	// loop, id, remove, add with my id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// TODO: Send a request when id is not found

}

func DeleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	// loops, id, remove (index, index+1)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// TODO: send a confirm or deny response
			break
		}
	}

}
