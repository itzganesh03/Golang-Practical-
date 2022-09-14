package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to class of mymodules in Golang")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	// http.ListenAndServe(":4000", r)
	log.Fatal(http.ListenAndServe(":4000", r))

}

func greeter() {
	fmt.Println("Hey there mod users")

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to Golang Class</h1>"))

}
