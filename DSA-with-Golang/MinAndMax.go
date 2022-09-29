package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{0, 1, 2, 2, -1}
	sort.Ints(a)
	fmt.Println("Minimum:", a[0])
	fmt.Println("Maximum:", a[len(a)-1])
}
