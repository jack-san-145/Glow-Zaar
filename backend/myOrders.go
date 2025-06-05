package main

import (
	"fmt"
	"net/http"
)

func getMyOrders(w http.ResponseWriter, r *http.Request) {
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
		fmt.Println("cookie - ", cookie)
	}

}
