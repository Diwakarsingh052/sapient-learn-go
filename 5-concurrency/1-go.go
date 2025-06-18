package main

import (
	"fmt"
	"time"
)

// concurrency is dealing with a lot of things at once
// parallelism is doing multiple things at once

func main() {
	// main func job is just to spin up the goroutine, not to wait for it
	go hello()
	time.Sleep(time.Second * 1) // this is very bad practice
}

func hello() {
	fmt.Println("hello")
}
