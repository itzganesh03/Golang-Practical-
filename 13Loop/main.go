package main

import "fmt"

func main() {
	fmt.Println("Welcome to class of Loop in Golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Firday", "Saturday"}

	// fmt.Println(days)

	// Normal Style Loop Method..........

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// Golang Short Method...............

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// Golang index Loop Method......
	for index, day := range days {
		fmt.Printf("index value is %v and value is %v\n", index, day)
	}

	// Golang While Loop..........

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto ganesh

		}
		if rougueValue == 5 {
			break
		}
		fmt.Println("value is:", rougueValue)
		rougueValue++
	}
ganesh:
	fmt.Println("This is Goto Statement")
}
