package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// implicit type and No var add

	welcome := "welcome to Pizza-Hut"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the  rating for our Pizza")

	// Comma ok syntex Pepole also called is Error ohk sysntx

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating,", input)
	fmt.Printf("Type of rating is: %T", input)

}
