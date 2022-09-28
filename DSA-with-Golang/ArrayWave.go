package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{12, 22, 32, 44, 55, 66, 77, 88, 93}
	sort.Ints(a)
	var b []int
	for i:=0;i<(len(a)-1)/2;i++{
		b=append(b, a[i],a[len(a)-i-1])
	}
	fmt.Println("Array Wave:",b)
}
