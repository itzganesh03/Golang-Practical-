package main

import (
	"fmt"
)

func main() {
	a := "MMMDCCXXIV"
	i := 0
	b := 0
	t := ""
	for c := 0; c <= len(a)-1; c++ {
		switch string(a[c]) {
		case "I":
			if t == "V" || t == "X" || t == "L" || t == "C" || t == "D" || t == "M" {
				i = i + b
				b = 0
			}
			b = b + 1
			t = "I"
		case "V":
			if t == "X" || t == "L" || t == "C" || t == "D" || t == "M" {
				i = i + b
				b = 0
			} else if t == "I" {
				i = i - b
				b = 0
			}
			b = b + 5
			t = "V"
		case "X":
			if t == "L" || t == "C" || t == "D" || t == "M" {
				i = i + b
				b = 0
			} else if t == "I" || t == "V" {
				i = i - b
				b = 0
			}
			b = b + 10
			t = "X"
		case "L":
			if t == "C" || t == "D" || t == "M" {
				i = i + b
				b = 0
			} else if t == "I" || t == "V" || t == "X" {
				i = i - b
				b = 0
			}
			b = b + 50
			t = "L"
		case "C":
			if t == "D" || t == "M" {
				i = i + b
				b = 0
			} else if t == "I" || t == "V" || t == "X" || t == "L" {
				i = i - b
				b = 0
			}
			b = b + 100
			t = "C"
		case "D":
			if t == "M" {
				i = i + b
				b = 0
			} else if t == "I" || t == "V" || t == "X" || t == "L" || t == "C" {
				i = i - b
				b = 0
			}
			b = b + 500
			t = "D"
		case "M":
			if t == "I" || t == "V" || t == "X" || t == "L" || t == "C" || t == "D" {
				i = i - b
				b = 0
			}
			b = b + 1000
			t = "M"
		}
		fmt.Println(string(a[c]), i, b)
	}
	i = i + b
	fmt.Println(i)
}
