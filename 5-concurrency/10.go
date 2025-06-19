package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go func() {
		for {
			time.Sleep(2 * time.Second)
			fmt.Println("waiting")
		}
	}()

	ch <- 10

}
