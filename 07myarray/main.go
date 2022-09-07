package main

import "fmt"

func main() {

	fmt.Println("Welcome to Array in Golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banna"
	fruitList[3] = "Peach"
	fmt.Println("Fruit List: ", fruitList)
	fmt.Println("Fruit List: ", len(fruitList))

	var vegyList = [3]string{"Tomato", "Potato", "Beans"}

	fmt.Println("Vegy List: ", vegyList)
	fmt.Println("Vegy List: ", len(vegyList))

}
