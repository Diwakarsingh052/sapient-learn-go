package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func main() {
	// context -> cancellation signals
	//     		-> passing values in different layers of a http request

	// empty container
	// we need a container to hold the context object
	// Background returns a non-nil, empty Context.
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel() // clean up the resources taken up by the context
	// if we don't put cancel in defer timer would be immediately cancelled

	fmt.Println(Slow(ctx, "10"))

}

// context should be the first argument passed to function

func Slow(ctx context.Context, input string) (int, error) {

	//sql.DB{}.ExecContext()
	i, err := strconv.Atoi(input)

	if err != nil {
		return 0, nil
	}
	time.Sleep(time.Second * 2)
	select {
	case <-ctx.Done(): // done would receive signal in case the timer is over
		fmt.Println("time is up")
		return 0, ctx.Err()
	default: // select would unblock immediately if done has not received the value
	}
	return i, nil
}
