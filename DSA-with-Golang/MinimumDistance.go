package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 2, 3, -1}
	var k int
	fmt.Println("Enter the 1st number: ")
	fmt.Scanln(&k)
	var l int
	fmt.Println("Enter the 2nd number: ")
	fmt.Scanln(&l)
	var x []int
	var y []int
	for i := range a {
		if a[i] == k {
			x = append(x, i)
		}
		if a[i] == l {
			y = append(y, i)
		}
	}
	temp := 999999999999
	for i := range x {
		for j := range y {
			if x[i] > y[j] && temp > x[i]-y[j] {
				temp = x[i] - y[j]
			} else if x[i] < y[j] && temp > y[j]-x[i] {
				temp = y[j] - x[i]
			}
		}
	}
	if temp != 999999999999 {
		fmt.Println("Minimum Distance:", temp)
	} else {
		fmt.Println("No valid Minimum Distance exists")
	}
}
