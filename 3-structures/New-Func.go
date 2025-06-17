package main

import (
	"fmt"
	"learn-go/db"
)

func main() {
	c := db.NewConn("postgres")
	fmt.Println(c)

}
