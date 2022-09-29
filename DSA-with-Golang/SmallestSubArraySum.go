package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	sort.Ints(a)
	var k int
	fmt.Println("Enter the value of K: ")
	fmt.Scanln(&k)
	c := -1
	var d []int
	for j := 0; j <= len(a)-1; j++ {
		temp := 0
		var s []int
		for i := j; i <= len(a)-1; i++ {
			if temp+a[i] <= k {
				temp = temp + a[i]
				s = append(s, a[i])
			}
			if temp == k {
				c = 0
				if len(d) == 0 {
					d = s
				} else if len(s) < len(d) {
					d = s
				}
				break
			}
		}
	}
	if c == -1 {
		fmt.Println("No valid Sub Array found")
	} else {
		fmt.Println("Smallest Sub Array Sum equal to", k, "is", d)
	}
}
