package main

import "fmt"

func main() {
	defer fmt.Println("to Class of defer in Golang")
	defer fmt.Println("Four")
	defer fmt.Println("Three")
	defer fmt.Println("Two")
	defer fmt.Println("One")
	// Defer Work with Revers Flow Of code
	fmt.Println("Welcome")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
