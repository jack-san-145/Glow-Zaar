package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Authentication(email string, password string) string {
	var emailFromDb, passwordFromDb string
	query := "select user_email,user_password from User where user_email = ? "
	row := Db.QueryRow(query, email)
	err := row.Scan(&emailFromDb, &passwordFromDb)
	if err == sql.ErrNoRows {
		return "Email doesn't exist"
	} else if email == emailFromDb && password == passwordFromDb {
		return "exactly matched"
	} else {
		return "Invalid Email or Password"
	}
}
