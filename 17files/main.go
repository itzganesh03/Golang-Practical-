package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welecome to class in files in Golang")

	content := "This is Files Method in Golang"
	//Create Automatic file in folder and write to content
	file, err := os.Create("./mygolangfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is:", length)
	defer file.Close()
}
