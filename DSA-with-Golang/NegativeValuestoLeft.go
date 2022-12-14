package main

import "fmt"

func MajorityElement(a []int) map[int]int {
	// a:= []int{27,22,23,25,21,22}
	dict := make(map[int]int)
	for i := range a {
		dict[a[i]] = dict[a[i]] + 1
	}
	return dict
}
func main() {
	a := []int{0, -1, 0, 1, 0, 1, 1}
	d := make(map[int]int)
	d = MajorityElement(a)
	var s []int
	for j := range d {
		if j < 0 {
			for i := 0; i < d[j]; i++ {
				s = append(s, j)
			}
		}
	}
	for i := range a {
		if a[i] >= 0 {
			s = append(s, a[i])
		}
	}
	fmt.Println(a, "array Move Negative to the Left is", s)
}
