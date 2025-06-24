package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type key string

// go get github.com/google/uuid
// go mod tidy // download all the required deps // remove unused deps

// The provided key must be comparable and should not be of type string or
// any other built-in type to avoid collisions between packages using context.
// Users of WithValue should define their own types for keys.
type ctxKey string

const Key ctxKey = "reqIdKey"

func main() {
	http.HandleFunc("/home", RequestId(Hello))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}

func Hello(w http.ResponseWriter, r *http.Request) {
	//var a any = "hello"
	//a = 10
	//a = true
	ctx := r.Context()
	// type assertion
	// checking if value is received of certain type or not
	// ok would be true if value is present and of the correct type
	reqId, ok := ctx.Value(Key).(string)
	if !ok {
		reqId = "unknown"
	}
	//if reqId == nil {
	//	reqId = "unknown"
	//}

	fmt.Println("req ended with ", reqId)
	fmt.Fprintln(w, "hello")
}

func RequestId(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqId := uuid.NewString()

		fmt.Println("req started with ", reqId)

		ctx := r.Context()

		// creating a new copy of context with new key value pairs added
		ctx = context.WithValue(ctx, Key, reqId)

		//withContext would update the internal context of the request object
		//with the updated context with values or timeouts
		next(w, r.WithContext(ctx))
	}
}
