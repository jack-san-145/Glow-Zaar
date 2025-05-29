package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func ConnectDB() {
	var err error
	Db, err = sql.Open("mysql", "root@tcp(localhost:3306)/GlowZaar")
	if err != nil {
		fmt.Println("Database connection failure")
	} else {
		fmt.Println("database connected")
	}
}
