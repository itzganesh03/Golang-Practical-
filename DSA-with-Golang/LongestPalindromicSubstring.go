package main

import (
	"fmt"
)

func main() {
	a := "abcba0123210"
	var s []string
	for i := 0; i <= len(a)-2; i++ {
		var x []string
		for j := i; j <= len(a)-1; j++ {
			x = append(x, string(a[j]))
			c := 0
			for j := range x {
				if x[j] != x[len(x)-j-1] {
					c = -1
				}
			}
			if c == 0 {
				if len(x) > len(s) {
					s = x
				}
			}
		}
	}
	fmt.Println(s, "is the Longest Palindrome Substring")
}
