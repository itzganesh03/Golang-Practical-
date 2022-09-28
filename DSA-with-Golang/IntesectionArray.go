package main

import (
	"fmt"
)
func main() {
	a := []int{0,1,2,-1}
	b := []int{0,5,2,2,4}
	var s= []int{}
	for j:=0;j<=len(a)-1;j++{
		for i:=0;i<=len(b)-1;i++{
		if(a[j]==b[i]){
			s=append(s,a[j])
	}
}
}
fmt.Println("Intersection Array between the array:",s)
}