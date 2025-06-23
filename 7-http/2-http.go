package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/text", data)
	http.ListenAndServe(":8080", nil)
}

func data(w http.ResponseWriter, r *http.Request) {
	//w http.ResponseWriter, is used to write resp to the client
	// http.Request// anything user send us would be in the request struct
	time.Sleep(time.Second * 5)
	ctx := r.Context() // this method gives the context from the request object

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		fmt.Println("user has left")
		return
	default:
		//client is still connected
	}

	w.WriteHeader(http.StatusOK)
	//w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("hello world"))
}
