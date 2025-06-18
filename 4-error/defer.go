package main

import (
	"log"
	"os"
)

func main() {
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//fmt.Println(3)
	////panic("panic")
	//return
	//fmt.Println(4)

	f, err := os.Open("file.txt")
	if err != nil {
		log.Println(err)
		return
	}
	// defer gurantees exec in case of panic and return
	// it runs when func returns
	defer f.Close()
	defer func() {

	}()
	f.Name()

}
