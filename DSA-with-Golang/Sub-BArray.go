package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	var k, n int
	fmt.Println("Enter the number of elements: ")
	fmt.Scanln(&n)
	var b = []int{}
	for i := 0; i <= n-1; i++ {
		fmt.Println("Enter the number at Position", i, ": ")
		fmt.Scanln(&k)
		b = append(b, k)
	}
	c := -1
	for j := 0; j <= len(a)-1; j++ {
		var s= []int {}
		for m := j; m <= len(a)-1; m++ {
			s = append(s, a[m])
		if reflect.DeepEqual(b, s) {
			c = 0
			fmt.Println("Subset Match for", b, "found")
			break
		}
	}
}
	if c == -1 {
		fmt.Println("No valid Subset found")
	}
}
