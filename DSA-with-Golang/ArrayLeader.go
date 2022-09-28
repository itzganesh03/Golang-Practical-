package main
import "fmt"
func SumArray(a[]int)int{
	temp:=0
	for i:=0;i<=len(a)-1;i++{
			temp=temp+a[i]
	}
	return temp
}

func main() {
	a := []int{1,12,3,3,8}
	c:=0
	for i:=0;i<=len(a)-2;i++{
		if(a[i]>=SumArray(a[i+1:len(a)])){
			c=1
		fmt.Println("Leader is",a[i],"at position",i)
		}
	}
	if(c==0){
		fmt.Println("No Leader found")
	}
}