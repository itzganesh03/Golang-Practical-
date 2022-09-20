package main

import (
	"fmt"
	"sync"
)

//channel says i am only allow in value when somebody listining to me..and this is big Problem Stetament

func main() {
	fmt.Println("This is class of Channel in Golang")

	myCh := make(chan int, 2)
	waitgroup := &sync.WaitGroup{}

	//<- this erro use for put value in channels
	// myCh <- 5
	// fmt.Println(<-myCh)
	waitgroup.Add(2)

	// this <- symbol in function means this is Outside in a box
	// Receive Only
	go func(ch <-chan int, waitgroup *sync.WaitGroup) {
		val, isChannelOpen := <-myCh

		fmt.Println(isChannelOpen)
		fmt.Println(val)
		// fmt.Println(<-myCh)
		waitgroup.Done()
	}(myCh, waitgroup)

	// this <- symbol in function means this is Inside in a box
	// Send Only
	go func(ch chan int, waitgroup *sync.WaitGroup) {
		myCh <- 0
		close(myCh)
		// myCh <- 6
		waitgroup.Done()
	}(myCh, waitgroup)

	waitgroup.Wait()
}
