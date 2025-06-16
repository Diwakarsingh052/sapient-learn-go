package main

import "learn-go/slice"

func main() {
	// make preallocates a backing array
	// very helpful if we have idea about the number of request ,
	//we can create size of array according to that
	// in this case append doesn't have to allocate the memory each time

	x := make([]int, 0, 10) // make(type,len,cap)
	slice.Inspect("x", x)
	x = append(x, 1, 2, 3)
	slice.Inspect("x", x)

}
