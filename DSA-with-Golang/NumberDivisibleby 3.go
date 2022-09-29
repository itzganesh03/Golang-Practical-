package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 6, 4, 9}
	for i := range a {
		if a[i]%3 == 0 {
			fmt.Println("Found", a[i], "at index", i, "Divisible by 3")
		}
	}
}
