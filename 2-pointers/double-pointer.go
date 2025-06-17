package main

import "fmt"

func main() {

	var p *int // nil
	fmt.Println(&p)
	updateNilPointer(&p)
	fmt.Println(*p)

}

func updateNilPointer(p1 **int) {
	x := 10

	// trying to access the value of p1
	// which is also another pointer named as p from the main function
	*p1 = &x // it directly changes p from the main function itself
	//var i int = &x

}
