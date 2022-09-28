package main

import (
	"fmt"
)

func main() {
	a := []int{0,1,2,2,-1}
	c:=-1
	for j:=0;j<=len(a)-3;j++{
		temp:=0
		var s []int
		for i:=j;i<=j+2;i++{
			temp=temp+a[i]
			s = append(s,a[i])
		}
		if(temp==0){
			c=0
			fmt.Println("Triplet whose sum is 0:=",s)
	}
}
if(c==-1){
	fmt.Println("No valid Triplet found")
}
}