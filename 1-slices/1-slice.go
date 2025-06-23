package main

import "fmt"

// global
var s int

//var i = true

func main() {

	//var s string = "hello"
	//var b bool
	//var i int
	//var f float64
	//
	//s1 := "hello"
	//var s2 = "hello"

	//var s [4]int
	var s []int
	fmt.Printf("%#v\n", s)

	x := []int{1, 2, 3, 4, 5} // this can be used in the local scope only
	fmt.Println(x)

}
