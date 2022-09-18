package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var waitgroup sync.WaitGroup //pointer

var mutex sync.Mutex

func main() {
	fmt.Println("The class of Go-routine in Golang")
	// go greeter("hello")
	// greeter("world")
	websitelist := []string{
		"https://jio.com",
		"http://go.dev",
		"http://google.com",
		"http://facebook.com",
		"http://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		waitgroup.Add(1)
	}
	waitgroup.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}

// }

//Wait group in Golang

func getStatusCode(endpoint string) {
	defer waitgroup.Done()

	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mutex.Lock()
		signals = append(signals, endpoint)
		mutex.Unlock()

		fmt.Printf("%d Status code for %s\n", result.StatusCode, endpoint)
	}

}

// Mutex is a Mutual exclusion Lock
