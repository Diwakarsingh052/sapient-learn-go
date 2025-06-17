package main

import "fmt"

func main() {
	i := 10    // x80
	var p *int // p is a pointer, it stores the memory address of i
	p = &i
	fmt.Println(p)
	fmt.Println(*p) // derefrecning the memory // access value at the address
	fmt.Println(&i)
	update(p)
	fmt.Println(i)
}
func update(p *int) {
	*p = 100
	// x80 = 100

}
