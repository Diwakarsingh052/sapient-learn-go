package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50}

	// if we don't want to modify the orignal slice than we must do a copy using make,and copy
	s1 := make([]int, len(s), 1000)

	copy(s1, s)
	s1[0] = 999
	fmt.Println(s)
	fmt.Println(s1)
}
