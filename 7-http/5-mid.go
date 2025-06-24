package main

import (
	"fmt"
	"net/http"
)

// req -> mid1->mid-2-> handler->services
//
//	<-		<-		<-
func main() {
	http.HandleFunc("/home", Mid1(Mid(home)))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home handler")
	fmt.Fprintln(w, "home handler")

}

func Mid1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("mid1 layer started")
		fmt.Println("pre processing logic going on mid1")
		next(w, r)
		fmt.Println("post processing logic going on mid1")
		fmt.Println("mid1 layer finished")
	}
}

func Mid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("mid layer started")
		fmt.Println("pre processing logic going on")
		next(w, r)
		fmt.Println("post processing logic going on")
		fmt.Println("mid layer finished")
	}
}

func calc(f func(int)) {}

type op func(int)

func add() op {
	return func(x int) {
		fmt.Println(x)
	}
}
