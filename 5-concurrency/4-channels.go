package main

import (
	"fmt"
	"sync"
)

// https://go.dev/ref/spec#Send_statements

// A send on an unbuffered channel can proceed if a receiver is ready.
// send will block until there is no recv
// channels are only meant to be used in concurrent programming
func main() {
	wg := new(sync.WaitGroup)
	ch := make(chan int) // unbuffered chan // size is 0
	wg.Add(2)
	go func() {
		defer wg.Done()
		ch <- 10 //A send on an unbuffered channel can proceed if a receiver is ready.
	}() // send

	go func() {
		defer wg.Done()
		fmt.Println(<-ch) // recv // it is always blocking until the value is not received
	}()

	wg.Wait()

}
