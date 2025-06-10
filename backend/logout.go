package main

import (
	"glow/database"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session_id")
	if cookie != nil {
		database.DeleteSession(cookie.Value)
		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    "",
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
		})
	}
}
