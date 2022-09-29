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
	a := []int{1, 2, 1, 4, 5, 6, 2, 2}
	d := make(map[int]int)
	d = MajorityElement(a)
	temp := 0
	tmax := 0
	for i := range d {
		if d[i] > temp {
			tmax = i
			temp = d[i]
		}
	}
	fmt.Println("Max Occuring Element is", tmax, "and occurs", temp, "times")
}
