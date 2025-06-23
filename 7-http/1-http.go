package main

import (
	"fmt"
	"net/http"
)

func main() {

	// registering routes
	// func would be called when someone hits the endpoint
	http.HandleFunc("/home", Home)

	// starting up the service on 8080
	err := http.ListenAndServe(":8080", nil)
	fmt.Println("server started on port 8080")
	if err != nil {
		panic(err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	// by default everything would be running as concurrent

	// this panic will be automatically recovered by http service
	panic("something went wrong in home func")

	//go func() {
	//	// panic inside a goroutine would stop the service
	//	// we need to recover from this manually
	//	defer recoverFunc()
	//	panic("something went wrong")
	//}()

	fmt.Println("Hello, World!")
}

func recoverFunc() {
	msg := recover()
	if msg != nil {
		// if this condition is true then panic happened
		fmt.Println("recovered from the panic:")
		fmt.Println(msg)
		//fmt.Println(string(debug.Stack()))
	}
}
