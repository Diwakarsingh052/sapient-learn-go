package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	ch := make(chan int, 10)
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 50; i++ {
			ch <- i
		}
		close(ch) // // sends a signal to stop the range // sending signal without data
		// close signal range that no more values be sent and it can stop after receiving remaining values
		//ch <- 10  // once the channel is closed, we can't send more values to it
	}()

	go func() {
		defer wg.Done()
		for v := range ch { // it would run infinitely, channel needs to be closed to stop this range
			fmt.Println(v)
		}
		//fmt.Println("done", <-ch) // after channel is closed you can receive the value, it would give default value of the ype
	}()

	wg.Wait()
}
