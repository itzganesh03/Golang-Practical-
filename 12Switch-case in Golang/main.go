package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welecome to class Switch Case in Golang ")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("DiceNumber is:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can Open")
	case 2:
		fmt.Println("You can move to 2 Spot")
		fallthrough
	case 3:
		fmt.Println("You can move to 3 Spot")
		fallthrough
	case 4:
		fmt.Println("You can move to 4 Spot")
	case 5:
		fmt.Println("You can move to 5 Spot")
		fallthrough
	case 6:
		fmt.Println("You can move to 6 Spot and roll dice again")
	default:
		fmt.Println("What was that")
	}
}
