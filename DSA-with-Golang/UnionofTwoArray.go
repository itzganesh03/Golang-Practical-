package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 2, -1}
	b := []int{0, 5, 2, 2, 4}
	var s = []int{}
	for j := 0; j <= len(a)-1; j++ {
		c := 0
		for i := 0; i <= len(b)-1; i++ {
			if a[j] == b[i] {
				c = 1
			}
		}
		if c == 0 {
			s = append(s, a[j])
		}
	}
	for j := 0; j <= len(b)-1; j++ {
		c := 0
		for i := 0; i <= len(a)-1; i++ {
			if b[j] == a[i] {
				c = 1
			}
		}
		if c == 0 {
			s = append(s, b[j])
		}
	}
	fmt.Println("Union Array between the array:", s)
}
