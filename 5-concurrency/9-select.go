package main

import (
	"fmt"
	"sync"
)

func main() {

	// don't use this pattern with buffered channels
	get := make(chan string)
	post := make(chan string)
	put := make(chan string)
	done := make(chan struct{})
	wg := &sync.WaitGroup{}
	wgWorker := &sync.WaitGroup{}

	wgWorker.Add(1)
	go func() {
		defer wgWorker.Done()
		//time.Sleep(4 * time.Second)
		get <- "get"
		fmt.Println("sent get")
	}()

	wgWorker.Add(1)
	go func() {
		defer wgWorker.Done()

		//time.Sleep(2 * time.Second)
		post <- "post"
		fmt.Println("sent post")
	}()

	wgWorker.Add(1)
	go func() {
		defer wgWorker.Done()
		put <- "put"
		put <- "put 2"
		fmt.Println("sent put")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		wgWorker.Wait()
		close(done)
	}()

	//fmt.Println(<-get)
	//fmt.Println(<-post)
	//fmt.Println(<-put)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for { // i := 0; i < 3; i++

			select {

			//whichever case is not blocking exec that first
			//whichever case is ready first, exec that.
			// possible cases are chan recv , send , default

			case msg := <-get: // channel receive
				fmt.Println(msg)
			case msg := <-post:
				fmt.Println(msg)
			case msg := <-put:
				fmt.Println(msg)

			case <-done: // receive a signal from done, and we are storing the done value in a variable
				fmt.Println("all work has been done")
				return

			}

		}

	}()

	wg.Wait()

}
