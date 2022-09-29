package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 2, 2, -1}
	c := -1
	for j := 1; j <= len(a)-2; j++ {
		ls := 0
		rs := 0
		var l []int
		var r []int
		for i := range a[0 : j-1] {
			ls = ls + a[i]
			l = append(l, a[i])
		}
		for i := range a[j+1 : len(a)-1] {
			rs = rs + a[i]
			r = append(r, a[i])
		}
		if ls == rs {
			c = 0
			fmt.Println("Equilibrium Point is", j, ":=", l, r)
		}
	}
	if c == -1 {
		fmt.Println("No valid Points found")
	}
}
