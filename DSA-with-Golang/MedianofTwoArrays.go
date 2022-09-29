package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{0, 1, 2, -1}
	b := []int{0, 5, 2}
	var s = []int{}
	for j := 0; j <= len(a)-1; j++ {
		s = append(s, a[j])
	}
	for i := 0; i <= len(b)-1; i++ {
		s = append(s, b[i])
	}
	sort.Ints(s)
	fmt.Println("Median of", s, "Array is", s[(len(s)-1)/2])
}
