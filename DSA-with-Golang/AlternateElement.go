package main

import "fmt"

func main() {
	a := []int{12, 22, 32, 44, 55, 66, 77, 88, 93, 93}
	for i := range a {
		if i%2 == 0 {
			fmt.Println(a[i])
		}
	}
}
