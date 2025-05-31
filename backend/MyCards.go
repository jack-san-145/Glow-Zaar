package main

import (
	"fmt"
	"net/http"
)

var CookieFound = make(map[string]bool)

func GetMyCards(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("Cookie_5")
	if cookie != nil {
		fmt.Println("True running")
		CookieFound["isFound"] = true
	} else {
		fmt.Println("False running")
		CookieFound["isFound"] = false
	}
	WriteJson(w, http.StatusAccepted, r, CookieFound)
}
