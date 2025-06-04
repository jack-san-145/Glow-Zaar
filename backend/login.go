package main

import (
	"fmt"
	"glow/database"
	"net/http"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		WriteError(w, http.StatusBadRequest, " requesting method mismatch")
		return
	}
	r.ParseForm()
	email := r.FormValue("user_email")
	password := r.FormValue("user_password")
	fmt.Println("Email - ", email)
	fmt.Println("Password - ", password)
	msg, isValidUser := database.Authentication(email, password)
	if !isValidUser {
		WriteJson(w, http.StatusAccepted, r, msg)
		return
	}
	session_id := database.GenerateSessionId(email, password)
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    session_id,
		Path:     "/",
		Expires:  time.Now().Add(3 * time.Hour),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	fmt.Println("cookie set successfully")
	WriteJson(w, http.StatusAccepted, r, msg)

}
