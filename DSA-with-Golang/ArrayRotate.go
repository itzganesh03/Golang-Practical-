package main

import (
	"fmt"
)

func main() {
	a := []int{12, 22, 32, 44, 55, 66, 77, 88, 93}
	var k int
	fmt.Println("Enter the value of K: ")
	fmt.Scanln(&k)
	var b [] int
	for i:=len(a)-k-1;i<=len(a)-1;i++{
		b=append(b, a[i])
	}
	for i:=0;i<len(a)-k-1;i++{
		b=append(b, a[i])
	}
	fmt.Println("Array after rotation by",k,":= ",b)
}
