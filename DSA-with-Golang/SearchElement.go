package main
import "fmt"
func main() {
	a := []int{12, 22, 32, 44, 55, 66, 77, 88, 93, 93}
	var temp int
	fmt.Scanln(&temp)
	c:=-1
	for i:=range a{
		if(a[i]==temp){
		c=i
		break;
		}
	}
	if(c==-1){
		fmt.Println("Search Element not found")
	}else{
	fmt.Println("Search Element found at",c)
	}
}