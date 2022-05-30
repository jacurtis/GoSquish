package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var connection *sql.DB

func Init() {
	var err error
	connection, err = sql.Open("mysql", "root:@/squish")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	err = connection.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Println("Database Successfully Connected")
}
