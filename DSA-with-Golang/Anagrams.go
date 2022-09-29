package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []string{"a", "b", "c"}
	b := []string{"c", "a", "b"}
	sort.Strings(a)
	sort.Strings(b)
	c := -1
	for j := 0; j <= len(a)-1; j++ {
		if a[j] != b[j] {
			c = 0
			break
		}
	}
	if c == -1 {
		fmt.Println("Anagram")
	} else {
		fmt.Println("Not an Anagram")
	}
}
