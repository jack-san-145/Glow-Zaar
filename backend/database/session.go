package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func GenerateSessionId(email string, password string) string {

	deleteQuery := "delete from Sessions where expires_at < now() "
	_, deleteErr := Db.Exec(deleteQuery)
	if deleteErr != nil {
		fmt.Println("Error while deleting the sesions")
	}

	sessionId := uuid.New().String()
	query := "insert into Sessions(session_id,user_id,created_at,expires_at) values(?,?,now(),date_add(now(),interval 3 hour))"
	fmt.Println("sessionId - ", sessionId)
	_, err := Db.Exec(query, sessionId, FindUserId(email, password))
	if err != nil {
		fmt.Println("Error while inserting seesion id's ")
	}
	return sessionId
}

func DeleteSession(session_id string) {
	query := "delete from Sessions where session_id = ? "
	_, err := Db.Exec(query, session_id)
	if err != nil {
		fmt.Println("error while deleting session - ", err)
	}
}
