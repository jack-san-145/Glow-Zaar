package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Authentication(email string, password string) (map[string]string, bool) {
	var emailFromDb, passwordFromDb string
	query := "select user_email,user_password from User where user_email = ? "
	row := Db.QueryRow(query, email)
	err := row.Scan(&emailFromDb, &passwordFromDb)
	if err == sql.ErrNoRows {
		return map[string]string{"message": "Email doesn't exist", "status": "false"}, false
	} else if email == emailFromDb && password == passwordFromDb {
		return map[string]string{"message": "Logged in succcessfully", "status": "true"}, true
	} else {
		return map[string]string{"message": "Invalid Email or Password", "status": "false"}, false
	}
}
