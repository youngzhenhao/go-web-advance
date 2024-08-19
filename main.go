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

var db *sql.DB

func initMysql() error {
	// DSN:Data Source Name
	dsn := username + ":" + password + "@" + protocol + "(" + host + ":" + port + ")" + "/" + dbname
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := initMysql(); err != nil {
		fmt.Println(err)
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	fmt.Println("Successfully connected!")
}
