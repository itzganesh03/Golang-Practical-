package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var s []float64
	for i := 0; i <= 2; i++ {
		var k float64
		fmt.Println("Enter a number: ")
		fmt.Scanln(&k)
		s = append(s, k)
	}
	sort.Float64s(s)
	if math.Sqrt((s[0]*s[0])+(s[1]*s[1])) == s[2] {
		fmt.Println(s, "is a Pythagorean Triplet")
	} else {
		fmt.Println(s, "is not a Pythagorean Triplet")
	}
}
