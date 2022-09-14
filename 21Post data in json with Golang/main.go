package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("This is class of Web Verb in Golang")
	// PerformGetRequest()
	PerformPostJsonRequest()

}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)

	}
	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content length", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is:", byteCount)
	fmt.Println(responseString.String())
	// fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	// fack json payload

	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with Golang",
			"price": 0,
			"platform":"learncodeonline.in"
		}
	
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
