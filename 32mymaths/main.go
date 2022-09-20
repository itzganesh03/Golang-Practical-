package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {
	fmt.Println("This is Class of math,crypto and random Number")

	// var myNumberOne int = 2
	// var myNumberTwo float64 = 4

	// fmt.Println("The sum is no:", myNumberOne+int(myNumberTwo))

	// how  to get random number in Golang
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	//random from crypto

	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNumber)
}
