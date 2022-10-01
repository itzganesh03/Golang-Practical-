package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat/combin"
)

func main() {
	a := "abcd"
	var b []string
	for i := 0; i <= len(a)-1; i++ {
		b = append(b, string(a[i]))
	}
	cs := combin.Permutations(len(b), len(b))
	var c []int
	for _, c = range cs {
		for i := 0; i <= len(c)-1; i++ {
			fmt.Print(b[c[i]])
		}
		fmt.Println()
	}
	fmt.Println("Number of permutations for", a, "is", len(cs))
}
