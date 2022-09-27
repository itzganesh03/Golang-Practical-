package main

import "fmt"

func main() {
	a := [...]int{21, 22, 23, 24, 25}
	var x int
	fmt.Println("Enter the number to compare the floor of: ")
	fmt.Scanln(&x)
	var temp = -999999999999
	for i := 0; i <= len(a)-1; i++ {
		if a[i] > temp && a[i] < x {
			temp = a[i]
		}
	}
	fmt.Println("Number just smaller than x: ", temp)
}
