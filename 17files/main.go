package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welecome to class in files in Golang")

	content := "This is Files Method in Golang"
	//Create Automatic file in folder and write to content
	file, err := os.Create("./mygolangfile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is:", length)
	defer file.Close()
	readFile("./mygolangfile.txt")
}

// read method in txt file
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Text data inside file is \n:", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
