package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("This class of Race condition")

	watigroup := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	mut.RLock()
	var score = []int{0}
	mut.RUnlock()

	// add three go function in Waitgroup...
	watigroup.Add(3)

	go func(waitgroup *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One routin")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		waitgroup.Done()
	}(watigroup, mut)

	go func(waitgroup *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two routin")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		waitgroup.Done()
	}(watigroup, mut)

	go func(waitgroup *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three routin")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		waitgroup.Done()
	}(watigroup, mut)

	// Always Use of this types of MUTEX lock don't Lock the whole resource......

	go func(waitgroup *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Four routin")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		waitgroup.Done()
	}(watigroup, mut)

	watigroup.Wait()
	fmt.Println(score)
}
