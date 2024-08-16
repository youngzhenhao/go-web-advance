package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "12345678"
	protocol = "tcp"
	host     = "127.0.0.1"
	port     = "3306"
	dbname   = "testdb"
)

func main() {
	// DSN:Data Source Name
	dsn := username + ":" + password + "@" + protocol + "(" + host + ":" + port + ")" + "/" + dbname
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
}
