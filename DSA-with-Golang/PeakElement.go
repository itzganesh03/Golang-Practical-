package main
import "fmt"
func main() {
	a := []int{12, 22, 32, 44, 55, 66, 77, 88, 93, 93}
	var temp=-99999999999999
	for i:=range a{
		if(a[i]>temp){
		temp=a[i]
		}
	}
	fmt.Println("Peak Element: ",temp)
}