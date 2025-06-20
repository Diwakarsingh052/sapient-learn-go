package main

import (
	"fmt"
	"sync"
)

var user = map[int]string{
	1: "John",
}
var x int = 1

// go run -race 11-mutex.go // run your program with race detector to check for data races

func main() {
	m := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go updateX(2, m, wg)
	go updateX(3, m, wg)

	fmt.Println("end of the main")
	wg.Wait()
}

func updateX(val int, m *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	// critical section
	// this is the place where we access the shared resource

	// when a goroutine acquires a lock, another goroutine can't access the critical section
	// until the lock is not released

	// data race situations
	//	- at least one concurrent write and n number of reads
	//	- n number of concurrent writes
	// 	- n number of concurrent writes and n number of concurrent reads
	// 	Note - Data race doesn't happen if there are only concurrent reads

	//m.Lock() // other goroutines wait here while lock is acquired
	//defer m.Unlock()
	x = val
	fmt.Println(x)
}
