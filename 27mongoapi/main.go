package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/itzganesh03/Golang-Practical-/mongoapi/router"
)

func main() {
	fmt.Println("This class of MangoDbApi in Golang")
	r := router.Router()
	fmt.Println("Server is Getting Started....")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at Port 4000 .....")
}
