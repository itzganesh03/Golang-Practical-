package main

import (
	"fmt"
)

func main() {
	a := []int{12, 22, 32, 44, 55, 66, 77, 88, 93}
	c := 0
	for i := 0; i <= len(a)-2; i++ {
		for j := i; j <= len(a)-1; j++ {
			if a[i] > a[j] {
				c = 1
				break
			}
		}
		if c != 0 {
			break
		}
	}
	if c == 0 {
		fmt.Println(a, "is sorted")
	} else {
		fmt.Println(a, "is not sorted")
	}
}
