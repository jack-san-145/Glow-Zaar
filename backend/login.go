package main

import (
	"fmt"
	"glow/database"
	"net/http"
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
	msg := database.Authentication(email, password)
	fmt.Println("msg - ", msg)
}
