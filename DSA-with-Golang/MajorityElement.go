package main
import (
	"fmt"
)
func main(a int[]) {
	// a:= []int{27,22,23,25,21,22}
	dict:=make(map[int]int)
	for i:= range a{
		dict[a[i]]=dict[a[i]]+1
	}
	fmt.Println(dict)
	return dict
}