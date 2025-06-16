package main

import "learn-go/slice"

// https://go.dev/ref/spec#Appending_and_copying_slices
func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70}
	b := x[1:6]

	slice.Inspect("x", x)
	slice.Inspect("b", b)

	b = append(b, 777, 888) // this would create a new backing array for b,

	slice.Inspect("x", x)
	slice.Inspect("b", b)

	b = append(b, 999)
	slice.Inspect("x", x)
	slice.Inspect("b", b)

}
