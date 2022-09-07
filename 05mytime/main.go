package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	// without Giving any format
	fmt.Println(presentTime)
	// Giving Specific Format
	fmt.Println(presentTime.Format("02-01-2006 Monday"))

	createdDate := time.Date(1999, time.May, 03, 23, 28, 0, 0, time.UTC)
	fmt.Println(createdDate)
	// 	fmt.Println(createdDate.Format("05-06-2020 Monday"))
	fmt.Println(createdDate.Format("02/01/2006 Friday"))
}
