package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{12, 22, 32, 44, 55, 66, 77, 88, 93}
	sort.Ints(a)
	fmt.Println("The 2nd highest element is ", a[len(a)-2])
}
