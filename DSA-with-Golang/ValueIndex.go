package main
import "fmt"
func main() {
	a := []int{1,2,3,3,4}
	for i:=range a{
		if(a[i]==i){
		fmt.Println("Found",i,"at index",a[i])
		}
	}
}