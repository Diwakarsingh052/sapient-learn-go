package main

import (
	"fmt"
	"sync"
)

// https://go.dev/ref/spec#Send_statements
// A send on a buffered channel can proceed if there is room in the buffer.

func main() {
	wg := new(sync.WaitGroup)
	ch := make(chan int, 5) // buffered channel

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}()

	go func() {
		defer wg.Done()

		for i := 1; i <= 10; i++ {
			abc := <-ch
			fmt.Println(abc)
		}
	}()

	wg.Wait()
}
