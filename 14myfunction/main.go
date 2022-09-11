package main

import "fmt"

func main() {
	fmt.Println("Welcome to Class of function in Golang")
	ganesh()

	// important Note you can not Declare Function in function......
	// Only allowed Outside of Fucntion.......

	result := adder(5, 5)
	fmt.Println("Result is :", result)

	proRes, myMessage := proAdder(2, 5, 8, 7, 9)
	fmt.Println("Pro Result is:", proRes)
	fmt.Println("Pro Meassage is:", myMessage)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	// Also  we have to Pass string in Return method
	return total, "hi this is ProAdder"
}

func ganesh() {
	fmt.Println("This is Custom Function")
	fmt.Println("hello i am Ganesh")
}
