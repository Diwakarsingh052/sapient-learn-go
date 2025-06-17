package main

import "fmt"

var x int = 10

func main() {
	var p *int // nil // default value of a pointer is nil
	fmt.Println(&p)
	// after calling the update value p is still nil, as we never updated the pointer
	updateValue(p)
	//fmt.Println(*p) // this line would panic,

}

func updateValue(p1 *int) {

	p1 = &x
	//fmt.Println(*p1)
}
