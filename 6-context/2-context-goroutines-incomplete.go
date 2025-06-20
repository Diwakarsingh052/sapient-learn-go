package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := slowFunc()
		ch <- x
	}()

	select {
	case x := <-ch:
		fmt.Println(x)
	case <-ctx.Done():
		fmt.Println("time is up")
	}
	wg.Wait()
	fmt.Println("end of the main")
}

func slowFunc() int {
	time.Sleep(time.Second * 3)
	fmt.Println("slowFunction: slow fn ran and add 100 records to db")
	fmt.Println()
	return 100
}
