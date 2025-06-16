package main

import (
	"learn-go/slice"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	//a[5] = 100 // update // not add
	slice.Inspect("a", a)
	a = append(a, 6)
	slice.Inspect("a", a)

	a = append(a, 7)
	slice.Inspect("a", a)
}
