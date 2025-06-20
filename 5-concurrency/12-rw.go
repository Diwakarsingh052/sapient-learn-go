package main

import (
	"fmt"
	"sync"
	"time"
)

//	type invoice struct {
//		name string
//		seat string
//		time time.Time
//	}
type Theater struct {
	seats   int          // Total seats available in the theater
	invoice chan string  // Channel to store the name of the person whose seat is booked
	rw      sync.RWMutex // Read-write Mutex for synchronizing shared access to Seats
}

var wg = new(sync.WaitGroup)
var wgBook = new(sync.WaitGroup)

func (t *Theater) BookSeat(name string) {
	defer wgBook.Done()
	// when Write lock is acquired, no other read or writes are allowed
	t.rw.Lock()
	// Releases the write lock when func completes
	defer t.rw.Unlock()
	if t.seats > 0 {
		// Simulate a seat booking-making process
		fmt.Println("Seat is available for", name)
		time.Sleep(2 * time.Second)
		fmt.Println("Booking confirmed", name)
		t.seats--
		t.invoice <- name
	} else {
		fmt.Println("No seats available", name)
	}

}

// checkSeats method checks the available Seats in the Theater
func (t *Theater) checkSeats() {
	defer wg.Done() // Decrement the counter when checkSeats is done executing
	// Acquire a lock for reading
	t.rw.RLock()
	// Releases the read lock when func completes

	//no one can write when read lock is acquired,
	// there could be unlimited number of reads
	defer t.rw.RUnlock()
	fmt.Println("Available Seats:", t.seats)
}

// printInvoice method prints the invoice for all booked seats
func (t *Theater) printInvoice() {
	defer wg.Done() // Decrement the counter when func is done executing

	for name := range t.invoice {
		fmt.Printf("Invoice is sent to %s\n", name)
	}
}

func main() {

	t := Theater{
		seats:   2,
		invoice: make(chan string),
		rw:      sync.RWMutex{},
	}

	// Start checkSeat routines
	for i := 1; i <= 6; i++ {
		wg.Add(1) // Increment wait group counter
		go t.checkSeats()
	}

	for i := 1; i <= 6; i++ {
		wgBook.Add(1)
		go t.BookSeat(fmt.Sprintf("Person %d", i))
	}

	// Start checkSeat routines
	for i := 1; i <= 6; i++ {
		wg.Add(1) // Increment wait group counter
		go t.checkSeats()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		wgBook.Wait()
		close(t.invoice)
	}()

	wg.Add(1)
	go t.printInvoice()

	wg.Wait()

}
