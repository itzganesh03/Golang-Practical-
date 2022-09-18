package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This Class of Go-Routines")
	go greeter("hello")
	greeter("world")

}

func greeter(s string) {

	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)

	}

}
