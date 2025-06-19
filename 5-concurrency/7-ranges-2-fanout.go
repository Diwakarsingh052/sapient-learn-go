package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wgWorker := &sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 1000000; i++ {
			//we need to block our goroutine before closing the channel
			//because we want to make sure all the work
			// is done and finished
			// wgWorker waitgroup we are using to track number of worker goroutines

			wgWorker.Add(1)
			go func(id int) {
				defer wgWorker.Done()
				ch <- id

			}(i)

		}
		// waiting until all the workers are not finished
		wgWorker.Wait()
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("received", v)
		}
	}()

	wg.Wait()
}
