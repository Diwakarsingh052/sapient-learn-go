package main

import (
	"fmt"
)

type Student struct {
	name string
	age  int
}

// method signature
//func (receiver) funcSignature {//body}

func (s *Student) SayHello() {
	fmt.Println("Hello, my name is", s.name)
}

func (s *Student) updateName(name string) {
	s.name = name
}

// this is not allowed
//func (s Student) SayHello() string {
//	fmt.Println("Hello, my name is", s.name)
//}

func main() {

	s := Student{
		name: "Jim",
	}
	// methods could be only called using struct variable
	s.SayHello()
	fmt.Println(s)
	s.updateName("Bob")

	s.SayHello()

	//log.Logger{}.Println()

}
