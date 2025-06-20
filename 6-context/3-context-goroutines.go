package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 2)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := slowFuncV2()

		select {
		case ch <- x:
			fmt.Println("sent to channel")
		case <-ctx.Done():
			fmt.Println("time is up for sender to send the value")
			fmt.Println("undoing all the changes if required")
			return
		}

	}()

	select {
	case x := <-ch:
		fmt.Println(x)
	case <-ctx.Done():
		fmt.Println("time is up in main")
	}
	wg.Wait()
	fmt.Println("end of the main")
}

func slowFuncV2() int {
	time.Sleep(time.Second * 1)
	fmt.Println("slowFunction: slow fn ran and add 100 records to db")
	fmt.Println()
	return 100
}
