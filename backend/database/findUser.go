package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func FindUserId(email string, password string) int {
	var user_id int
	query := "select user_id from User where user_email = ? and user_password = ?"
	err := Db.QueryRow(query, email, password).Scan(&user_id)
	if err != nil {
		fmt.Println("error during finding the user_id")
	}
	return user_id
}

func FindUserBySessionId(session_id string) int {
	var UserID int
	query := "select user_id from User where session_id = ?"
	err := Db.QueryRow(query, session_id).Scan(&UserID)
	if err != nil {
		fmt.Println("error while finding userid from session id")
		return 0
	}
	return UserID
}
