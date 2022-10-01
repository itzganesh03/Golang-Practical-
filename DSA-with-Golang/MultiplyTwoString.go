package main

import (
	"fmt"
	"strconv"
)

func main() {
	a, e := strconv.Atoi("192")
	b, e := strconv.Atoi("2")
	_ = e
	fmt.Println(a * b)
}
