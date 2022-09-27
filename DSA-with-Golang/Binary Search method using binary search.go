package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 5, 10, 15, 20, 25, 30}
	sort.Ints(a)
	var min = 0
	var max = len(a) - 1
	var search int
	fmt.Println("Enter the number to be searched:")
	fmt.Scanln(&search)
	var mid = (min + max) / 2
	for {
		if min == max || mid == max || mid == min {
			fmt.Println("Search Value not found")
			break
		}
		if search > a[mid] {
			min = mid
		} else if search < a[mid] {
			max = mid
		} else if a[mid] == search {
			fmt.Println("Search value found", mid)
			break
		}
		mid = (min + max) / 2
	}
}
