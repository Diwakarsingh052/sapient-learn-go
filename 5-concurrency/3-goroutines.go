package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go work(i, wg)
	}

	wg.Wait()
}

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("doing work", id)
}
