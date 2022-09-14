package main

import (
	"encoding/json"
	"fmt"
)

// create Structure
// For back tik `` is for Created aliases like below Program
// For dash tik - is use for don't wont to reflect this filed who ever consume my json data
// omitempty use for just case value is Nill & Null so don't through at all

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"Web-site"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	// fmt.Println("This is Class of JSON Data in Golang")
	// EncodeJson()
	DecodeJson()
}

// Encoding the Json Data

func EncodeJson() {

	myCourses := []course{
		{"ReactJs BootCamp", 300, "dev.go.in", "REJ123", []string{"web-dev", "js"}},
		{"MERN BootCamp", 220, "dev.go.in", "MERN123", []string{"Full-stack", "js"}},
		{"Angular BootCamp", 399, "dev.go.in", "GNS123", nil},
	}

	// Package this data as JSON data
	// Marshal is the way to impliment of Json Data

	finalJson, err := json.MarshalIndent(myCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "MERN BootCamp",
		"Price": 220,
		"Web-site": "dev.go.in",
		"tags": ["Full-stack", "js"] 
		
	}
	`)
	//when we do coma after tags in jsondata then json is not valid. other wise it is valid
	var myCourses course

	checkvalid := json.Valid(jsonDataFromWeb)

	if checkvalid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &myCourses)
		fmt.Printf("%#v\n", myCourses)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and Type is %T\n", k, v, v)
	}
}
