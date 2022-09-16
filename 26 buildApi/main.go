// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/gorilla/mux"
// )

// // model of course -file

// type Course struct {
// 	CourseId    string  `json:"courseid"`
// 	CourseName  string  `json:"coursename"`
// 	CoursePrice int     `json:"price"`
// 	Author      *Author `json:"author"`
// }

// type Author struct {
// 	FullName string `json:"fullname"`
// 	Website  string `json:"website"`
// }

// // fake DB
// var courses []Course

// // middleware, helper -file

// func (c *Course) IsEmpty() bool {
// 	// return c.CourseId == "" && c.CourseName == ""
// 	return c.CourseId == ""

// }

// func main() {
// 	fmt.Println("This is class of BuildApi in Golang")
// 	// code in Mux router
// 	r := mux.NewRouter()

// 	// seeding the Data
// 	courses = append(courses, Course{CourseId: "2", CourseName: "React.js",
// 		CoursePrice: 300, Author: &Author{FullName: "Ganesh", Website: "ganesh.in.com"}})

// 	courses = append(courses, Course{CourseId: "4", CourseName: "Mearn stack",
// 		CoursePrice: 600, Author: &Author{FullName: "Divyani", Website: "divyani.in.com"}})

// 	// routing

// 	r.HandleFunc("/", serveHome).Methods("GET")
// 	r.HandleFunc("/courses", getAllCourses).Methods("GET")
// 	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
// 	r.HandleFunc("/course", createOneCourse).Methods("POST")
// 	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
// 	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

// 	// listen to a Port
// 	log.Fatal(http.ListenAndServe(":4000", r))
// }

// // controllers -file

// // serve home route

// func serveHome(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("<h1>Welcome to API by Golang</h1>"))
// }

// func getAllCourses(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Get All Courses")
// 	w.Header().Set("Content-Type", "applicatioan/json")
// 	json.NewEncoder(w).Encode(courses)
// }

// func getOneCourse(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Get One Course")
// 	w.Header().Set("Content-Type", "applicatioan/json")

// 	// i need to grab id from request
// 	params := mux.Vars(r)

// 	// fmt.Println("For know to Params Work:", params)

// 	// loop through courses, find matching id and return the response

// 	for _, course := range courses {
// 		if course.CourseId == params["id"] {
// 			json.NewEncoder(w).Encode(course)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode("No Course Found with Given Id")
// 	return
// }

// func createOneCourse(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Create One Course")
// 	w.Header().Set("Content-Type", "applicatioan/json")

// 	// what if mode is body is empty
// 	if r.Body == nil {
// 		json.NewEncoder(w).Encode("Please Send Some Data")
// 	}

// 	// what about data -{} send like this
// 	var course Course
// 	_ = json.NewDecoder(r.Body).Decode(&course)
// 	if course.IsEmpty() {
// 		json.NewEncoder(w).Encode("No Data is inside JSON")
// 		return
// 	}

// 	// generate unique id, and converted to string
// 	// append new course to courses

// 	// rand.Seed(time.Now().UnixNano())
// 	rand.Seed(time.Now().UnixNano())
// 	course.CourseId = strconv.Itoa(rand.Intn(100))
// 	courses = append(courses, course)
// 	json.NewEncoder(w).Encode(course)
// 	return

// }

// func updateOneCourse(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Create one Course:")
// 	w.Header().Set("content Type", "applicatioan/json")

// 	// first - grab id from request

// 	params := mux.Vars(r)

// 	// loops se value layenge after we Get id we go head remove opration and add with my id

// 	for index, course := range courses {
// 		if course.CourseId == params["id"] {
// 			courses = append(courses[:index], courses[index+1:]...)
// 			var course Course
// 			_ = json.NewDecoder(r.Body).Decode(&course)
// 			course.CourseId = params["id"]
// 			courses = append(courses, course)
// 			json.NewEncoder(w).Encode(course)
// 			return
// 		}
// 	}
// 	// Send a response when id is not found

// }

// func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Delete One Course")
// 	w.Header().Set("content Type", "applicatioan/json")

// 	params := mux.Vars(r)

// 	// First we go loop after id and perform remove opration and after take care of Index , index+1

// 	for index, course := range courses {
// 		if course.CourseId == params["id"] {
// 			courses = append(courses[:index], courses[index+1:]...)
// 			break
// 		}
// 	}

// }

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API - LearnCodeOnline.in")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Hitesh Choudhary", Website: "go.dev"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by LearnCodeOnline</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "applicatioan/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//TODO: check only if title is duplicate
	// loop, title matches with course.coursename, JSON

	// generate unique id, string
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop, id, remove, add with my ID

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
	//TODO: send a response when id is not found
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	params := mux.Vars(r)

	//loop, id, remove (index, index+1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// TODO: send a confirm or deny response
			break
		}
	}
}
