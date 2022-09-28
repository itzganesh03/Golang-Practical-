package main

import (
	"fmt"
	"math"
)

func main() {
	var k float64
	fmt.Println("Enter the value of K: ")
	fmt.Scanln(&k)
	fmt.Println("Square Root of",k,":= ",math.Sqrt(k))
}
