package main

import "fmt"

func main() {

	fmt.Println("Welecome to Class of Structc in Golang")

	//There is no inheritans in Golang
	// Also No super and Parent Method in Golang

	ganesh := User{"Ganesh", "ganesh@go.dev", true, 21}
	fmt.Println(ganesh)
	fmt.Printf("User Detials: %+v\n", ganesh)
	fmt.Printf("Name is %v and Email is %v", ganesh.Name, ganesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
