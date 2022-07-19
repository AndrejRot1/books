package main

import (
	"database/sql"
	"fmt"
	
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func connectDatabse() {
	db, err = sql.Open("mysql", "testdb2:eJ0Wya41pvHfZdVB_@tcp(sandbox.njorka.net:9836)/testdb2")
	fmt.Println("Database connected.")
}
