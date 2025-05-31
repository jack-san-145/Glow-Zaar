package main

import (
	"fmt"
	"glow/database"
	"net/http"
	"strconv"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		WriteError(w, http.StatusBadRequest, "Request method mismatch")
		return
	}
	r.ParseForm()
	name := r.FormValue("user_name")
	email := r.FormValue("user_email")
	password := r.FormValue("user_password")
	age := r.FormValue("user_age")
	address := r.FormValue("user_address")

	age_int, _ := strconv.Atoi(age)
	fmt.Println(name, email, password, age, address)
	msg, err := database.NewUserRegistration(name, email, password, age_int, address)
	fmt.Println("msg - ", msg)
	if err != nil {
		fmt.Println(err.Error())
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	WriteJson(w, http.StatusAccepted, r, msg)
}
