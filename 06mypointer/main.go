package main

import "fmt"

func main() {
	fmt.Println("Welcome to class of pointer in Golang")

	//******Created to Pointere***********

	// var ptr *int

	// fmt.Println("This is value of Pointer:", ptr)

	//----------------------------------------------------------
	// important Notes
	//----------------------------------------------------------

	// ********Created to Poiter and Refrence to some memory************
	// ********when we talk About Refrence we can use "&" this Symbol*********** refrence means this symbol &

	mynumber := 24

	var ptr = &mynumber

	fmt.Println("pointer address", ptr)
	fmt.Println("pointer value", *ptr)

	// addition in pointer value
	*ptr = *ptr + 2
	fmt.Println("New value is", mynumber)
}
