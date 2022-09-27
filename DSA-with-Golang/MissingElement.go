package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{27, 22, 23, 25, 21}
	sort.Ints(a)
	var temp = a[0]
	for i := 1; i <= len(a)-1; i++ {
		if temp+1 != a[i] {
			fmt.Println("Missing Element: ", temp+1)
		}
		temp = a[i]
	}
}
