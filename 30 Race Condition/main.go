package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("This class of Race condition")

	watigroup := &sync.WaitGroup{}
	var score = []int{0}

	watigroup.Add(3)
	go func(waitgroup *sync.WaitGroup) {
		fmt.Println("One routin")
		score = append(score, 1)
		waitgroup.Done()
	}(watigroup)

	go func(waitgroup *sync.WaitGroup) {
		fmt.Println("Two routin")
		score = append(score, 2)
		waitgroup.Done()
	}(watigroup)

	go func(waitgroup *sync.WaitGroup) {
		fmt.Println("Three routin")
		score = append(score, 3)
		waitgroup.Done()
	}(watigroup)

	watigroup.Wait()
}
