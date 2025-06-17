package main

import "fmt"

func main() {

	type Person struct {
		name    string // fields of struct
		age     int
		address string
	}

	//var loginDetails struct {
	//	username string
	//	password string
	//}

	var p Person
	p.name = "bob"

	fmt.Printf("%+v", p) // it can print field value pairs
}
