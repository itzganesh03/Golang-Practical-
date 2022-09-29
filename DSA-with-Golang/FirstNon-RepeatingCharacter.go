package main

import (
	"fmt"
)

func MajorityElement(a string) map[string]int {
	// a:= []int{27,22,23,25,21,22}
	dict := make(map[string]int)
	for i := range a {
		dict[string(a[i])] = dict[string(a[i])] + 1
	}
	return dict
}
func main() {
	a := "abacc"
	c := -1
	d := make(map[string]int)
	d = MajorityElement(a)
	for j := 0; j <= len(a)-1; j++ {
		if d[string(a[j])] == 1 {
			fmt.Println("First Non Repeating Character is", string(a[j]))
			c = 0
			break
		}
	}
	if c == -1 {
		fmt.Println("No Non Repeating Characters found")
	}
}
