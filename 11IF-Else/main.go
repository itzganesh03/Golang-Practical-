package main

import "fmt"

func main() {
	fmt.Println("Welcome to class of IF and Else in Golang")

	loginCount := 10

	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "watch out"
	} else {
		result = "Exactly 10 login Count"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Number is Less then 10")
	} else {
		fmt.Println("Number is Not Less then 10")
	}
}
