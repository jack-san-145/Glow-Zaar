package main

import (
	"fmt"
	"glow/database"
	"net/http"
)

var CookieFound = make(map[string]bool)

func GetMyCards(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session_id")
	if cookie != nil {
		fmt.Println("True running")
		CookieFound["isFound"] = true
		UserID = database.FindUserBySessionId(cookie.Value)
		WriteJson(w, http.StatusAccepted, r, CookieFound)
		fmt.Println("cookie - ", cookie)

	} else {
		fmt.Println("False running")
		CookieFound["isFound"] = false
		WriteJson(w, http.StatusAccepted, r, CookieFound)
		fmt.Println("cookie - ", cookie)
	}

}
