package main

import "fmt"

func main() {
	x := []int{10, 20, 30, 40, 50}
	updateList(x, 100)
	fmt.Println(x)

	x = appendToList(x, 200)
	fmt.Println(x)
}

func updateList(abc []int, i int) {
	// if you want to update the value ,
	//then passing the reference of existing backing array is fine
	abc[0] = i
}

func appendToList(abc []int, i int) []int {

	//fmt.Println("abc", append(abc, i))
	fmt.Println("abc 1", abc)
	return abc // make sure to return the updated slice if using append
}

// // if a function is not suppose to modify the original slice
//// then no update or append operation should be performed directly
