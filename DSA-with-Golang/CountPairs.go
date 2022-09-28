package main

import (
	"fmt"
)

func main() {
	a := []int{0,1,2,2,-1}
	c:=-1
	var k int
	fmt.Println("Enter the value of K: ")
	fmt.Scanln(&k)
	for j:=0;j<=len(a)-2;j++{
		temp:=0
		var s []int
		for i:=j;i<=j+1;i++{
			temp=temp+a[i]
			s = append(s,a[i])
		}
		if(temp==k){
			c=0
			fmt.Println("Pair whose sum is ",k,":=",s)
	}
}
if(c==-1){
	fmt.Println("No valid Pairs found")
}
}