package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB(database string) {
	var err error
	DB, err = sql.Open("mysql", database)
	if err != nil {
		fmt.Println("Failed Connect to Database: ", err)
		return
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println("Error Ping: ", err)
		return
	}
	fmt.Println("Connected to Database!")
}
