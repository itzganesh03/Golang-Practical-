package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to map class in Golang")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["RB"] = "ruby"
	languages["PY"] = "python"
	languages["Go"] = "golang"

	fmt.Println("List Of all languages", languages)
	//when if you want any single value of language map
	fmt.Println("Js short for:", languages["JS"])
	fmt.Println("PY short for:", languages["PY"])

	// when you delete any value of map list there is delete method in Golang

	delete(languages, "RB")
	fmt.Println("List of all languages:", languages)

	// loops are in Golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is  %v\n", key, value)
	}

}
