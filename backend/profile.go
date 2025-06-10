package main

import (
	"fmt"
	"glow/database"
	"glow/shared"
	"net/http"
)

func ProfileStatus(w http.ResponseWriter, r *http.Request) {
	var CookieFound = make(map[string]bool)
	cookie, _ := r.Cookie("session_id")
	if cookie != nil {
		fmt.Println("True running")
		CookieFound["isFound"] = true
		WriteJson(w, http.StatusAccepted, r, CookieFound)
		fmt.Println("cookie - ", cookie)

	} else {
		fmt.Println("False running")
		CookieFound["isFound"] = false
		WriteJson(w, http.StatusAccepted, r, CookieFound)
	}

}

func MyProfile(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session_id")
	if cookie != nil {
		UserID := database.FindUserBySessionId(cookie.Value)
		fmt.Println("User id profile - ", UserID)
		var profile shared.Profile
		profile = database.GetProfile(UserID)
		WriteJson(w, http.StatusAccepted, r, profile)
	}
}
