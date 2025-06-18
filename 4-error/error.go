package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	//var i []int
	//i[0] = 10

	// default value of error is nil
	i, err := strconv.Atoi("10")
	// error handling must be next line before proceeding with any kind of code
	if err != nil {
		log.Println(err)
		return
	}

	val, err := strconv.ParseInt("10", 10, 64)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(val)

	fmt.Println("some more important code that depends on i var", i)
}
