package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//variable Decleration

// const url = "http://dypiu-portalv2.epizy.com/"

const url = "https://lco.dev"

func main() {
	fmt.Println("This is class of webrequest in Golang")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type %T\n", response)

	defer response.Body.Close() //caller Responsibility to close the Connection

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
}
