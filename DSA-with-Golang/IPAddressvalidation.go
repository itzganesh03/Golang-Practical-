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
	a := "192.168.25.1"
	c := -1
	d := make(map[string]int)
	d = MajorityElement(a)
	if d["."] == 3 {
		for j := 0; j <= len(a)-1; j++ {
			if a[j] == 46 {
				continue
			} else if a[j] < 48 || a[j] > 57 {
				c = 0
				break
			}
		}
	} else {
		c = 0
	}
	if c == -1 {
		fmt.Println(a, "is a valid IP Address")
	} else {
		fmt.Println(a, "is not a valid IP Address")
	}
}
