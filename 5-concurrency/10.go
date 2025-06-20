package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//ch := make(chan int)
	//
	//go func() {
	//	for {
	//		time.Sleep(2 * time.Second)
	//		fmt.Println("waiting")
	//	}
	//}()
	//
	//ch <- 10
	do()

}

func do() {
	ch := make(chan int)
	wg := new(sync.WaitGroup)

	wg.Add(3)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for v := range ch {
			time.Sleep(10 * time.Millisecond)
			fmt.Println("g1", v)
		}
		fmt.Println("finished g1")

	}()

	go func() {
		defer wg.Done()
		for v := range ch {
			time.Sleep(10 * time.Millisecond)
			fmt.Println("g2", v)
		}
		fmt.Println("finished g2")
	}()

	wg.Wait()

}
