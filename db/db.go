package db

// var Db string = "postgres"

type Conn struct {
	db string
}

func NewConn(c string) Conn {
	return Conn{db: c}
}
