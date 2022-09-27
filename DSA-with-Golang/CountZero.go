package main

import "fmt"

func main() {
	a := [...]int{0, 1, 0, 1, 0, 1, 0}
	var count = 0
	for i := 0; i <= len(a)-1; i++ {
		if a[i] == 0 {
			count = count + 1
		}
	}
	fmt.Println("Count; ", count)

}
