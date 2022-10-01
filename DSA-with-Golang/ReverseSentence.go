package main

import (
	"fmt"
)

func main() {
	a := "This is Fun"
	var s []string
	var b string
	for i := 0; i <= len(a)-1; i++ {
		if string(a[i]) == " " {
			s = append(s, b)
			b = ""
		} else {
			b = b + string(a[i])
		}
	}
	s = append(s, b)
	b = ""
	for i := len(s) - 1; i >= 0; i-- {
		b = b + s[i] + " "
	}
	fmt.Println(b)
}
