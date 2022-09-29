package main

import (
	"fmt"
)

func main() {
	a := []int{10, 121, 1, 2, 1}
	// var y []int
	for i := range a {
		temp := a[i]
		var x []int
		for j := 0; j <= 100; j++ {
			x = append(x, temp%10)
			if temp/10 == 0 {
				break
			}
			temp = temp / 10
		}
		c := 0
		for j := range x {
			if x[j] != x[len(x)-j-1] {
				c = -1
			}
		}
		if c == 0 {
			fmt.Println(a[i], "is Palindrome")
		} else {
			fmt.Println(a[i], "is not Palindrome")
		}
	}
}
