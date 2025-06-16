package main

import "fmt"

// ptr,len, cap -> slice
func main() {
	a := []int{10, 20, 30, 40, 50}
	b := a[1:3] // [index:index+1]
	fmt.Println(b)
	b[0] = 999
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(len(a), len(b))

}
