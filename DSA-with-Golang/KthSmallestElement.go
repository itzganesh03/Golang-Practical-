package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{12, 22, 32, 44, 55, 66, 77, 88, 93}
	sort.Ints(a)
	var k int
	fmt.Println("Enter the value of K: ")
	fmt.Scanln(&k)
	if k > len(a)-1 || k < 1 {
		fmt.Println("Invalid value of K")

	} else {
		fmt.Println("The ", k, "th smallest element is ", a[k-1])
	}
}
