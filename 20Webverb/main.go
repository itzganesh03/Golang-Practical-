package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("This is class of Web Verb in Golang")
	PerformGetRequest()

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
