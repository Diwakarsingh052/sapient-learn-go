package main

import (
	"fmt"
	"runtime/debug"
)

func main() {

	DoSomething()

	fmt.Println("doing other imp tasks, that are not dependent on the update slice")
}

func DoSomething() {
	defer recoverFunc()
	updateSlice([]int{})
	fmt.Println("doing something func depends on the update slice")
}
func updateSlice(s []int) {

	s[0] = 100
	fmt.Println("update slice")
}

func recoverFunc() {
	msg := recover()
	if msg != nil {
		// if this condition is true then panic happened
		fmt.Println("recovered from the panic:")
		fmt.Println(msg)
		fmt.Println(string(debug.Stack()))
	}
}
