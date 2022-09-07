package main

import "fmt"

const Logintoken string = "This is Public Function Demo" //Public function Declearation....

func main() {
	var username string = "Ganesh"
	fmt.Println(username)
	fmt.Printf("this variable types is: %T \n", username)

	var isLoggedin bool = false
	fmt.Println(isLoggedin)
	fmt.Printf("this variable types is: %T \n", isLoggedin)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("this variable types is: %T \n", smallVal)

	var smallint int = 258
	fmt.Println(smallint)
	fmt.Printf("this variable types is: %T \n", smallint)

	var smallFlot32 float32 = 258.45445454545
	fmt.Println(smallFlot32)
	fmt.Printf("this variable types is: %T \n", smallFlot32)

	var smallFlot float64 = 258.45445454545
	fmt.Println(smallFlot)
	fmt.Printf("this variable types is: %T \n", smallFlot)

	// default values and aliases

	var anothervalue int
	fmt.Println(anothervalue)
	fmt.Printf("this variable types is: %T \n", anothervalue)

	// implicit type

	var website = "modiganesh83@gmail.com"
	fmt.Println(website)

	// No var style

	NumberofUser := 500000
	fmt.Println("This No var Style :", NumberofUser)

	fmt.Println(Logintoken)
	fmt.Printf("this variable type is: %T \n", Logintoken)

}
