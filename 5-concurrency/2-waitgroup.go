package main

import (
	"fmt"
	"sync"
)

func main() {
	// wg must be a pointer
	//wg := &sync.WaitGroup{}
	wg := new(sync.WaitGroup)

	// waitgroup counter represents number of goroutine we are running
	wg.Add(1)
	go func() {
		defer wg.Done() /// giving a guarantee that even
		// in case of panic this would decrement the counter
		fmt.Println("hello world")

	}()

	fmt.Println("some other code in main running")
	wg.Wait() // wait until the counter is not 0
	fmt.Println("end of the main")

}
