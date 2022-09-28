package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	var k, n int
	fmt.Println("Enter the 1st index to be swapped: ")
	fmt.Scanln(&n)
	var b = []int{}
	fmt.Println("Enter the 2nd index to be swapped: ")
	fmt.Scanln(&k)
	for j := 0; j <= len(a)-1; j++ {
		if(j==n-1){
		b = append(b, a[k-1])
		}else if(j==k-1){
			b = append(b, a[n-1])
		}else{
			b = append(b, a[j])
		}
	}
	fmt.Println(b)
}