package database

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

func NewUserRegistration(name string, email string, password string, age int, address string) (map[string]string, error) {
	var emailFromDb, passwordFromDb string
	query := "select user_email,user_password from User where user_email = ? "
	row := Db.QueryRow(query, email)
	err := row.Scan(&emailFromDb, &passwordFromDb)
	if err != sql.ErrNoRows {
		return map[string]string{"message": "Email already exist"}, nil
	}
	query = "insert into User(user_name,user_password,age,address,user_email) values(?,?,?,?,?)"
	_, err = Db.Exec(query, name, password, age, address, email)
	if err != nil {
		return nil, errors.New("error occured while inserting users records")
	}
	msg := " Successfully Registered '" + name + "' !!"
	return map[string]string{"message": msg}, nil
}
