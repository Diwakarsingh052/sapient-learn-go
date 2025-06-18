package main

import (
	"errors"
	"fmt"
)

var user = make(map[int]string, 100)

// variable which would be used to store errors should start with word Err

var ErrNotFound = errors.New("not found")

func main() {

	name, err := FetchRecord(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name)

}

// error must be the last value to be returned from function

func FetchRecord(id int) (string, error) {
	name, ok := user[id]
	if !ok {
		//return "", ErrNotFound
		return "", fmt.Errorf("user %d not found", id)
	}
	return name, nil
}
