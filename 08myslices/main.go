package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to Class of Slices in Golang")

	var fruitList = []string{"Apple", "Banana", "Peach", "Mango"}
	fmt.Println("Befor Append FruitList: ", fruitList)
	fmt.Println("FruitList: ", len(fruitList))
	fmt.Printf("Type of FruitList %T\n", fruitList)

	//Add value in array Use Append Method

	fruitList = append(fruitList, "Tomato", "Kiwi")
	fmt.Println("After Append FruitList:", fruitList)

	// Some of imp method of making Todo App and long values

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	// Memory Allocation Part

	// default allocation of Memory

	scorList := make([]int, 4)

	scorList[0] = 235
	scorList[1] = 436
	scorList[2] = 758
	scorList[3] = 928

	// When we use Append that time ReAllocation of Memory

	scorList = append(scorList, 658, 984)
	fmt.Println(scorList)

	// Here is Some intresting Method Called "Sort"

	sort.Ints(scorList)
	fmt.Println(scorList)
	// This is IntsAreSorted is Boolen mmethod Gives True and False
	fmt.Println(sort.IntsAreSorted(scorList))
}
