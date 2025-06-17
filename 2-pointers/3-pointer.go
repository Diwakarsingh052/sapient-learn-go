package main

import "fmt"

func main() {
	x := 10
	p := &x // var p *int = &x
	updateVal(p)

	fmt.Println(x)

}

func updateVal(px *int) {
	var randomVal = 200
	// we have changed the reference of px to store a new variable
	px = &randomVal // this should be avoided until have a clear important case

	// this would not update the main function x variable, it would update randomVal
	*px = 100

	fmt.Println("randomVal", *px)
}
