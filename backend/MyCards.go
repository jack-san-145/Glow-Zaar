package main

import (
	"fmt"
	"glow/database"
	"net/http"
)

func GetMyCardStatus(w http.ResponseWriter, r *http.Request) {
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

func GetMyCardProducts(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session_id")
	if cookie != nil {
		UserID := database.FindUserBySessionId(cookie.Value)
		fmt.Println("User id from mycart - ", UserID)
		CartProducts, err := database.DisplayCart(UserID)
		if err != nil {
			WriteError(w, http.StatusBadRequest, err.Error())

		} else {
			fmt.Println("cart products to send to frontend  - ", CartProducts)
			WriteJson(w, http.StatusAccepted, r, CartProducts)
		}
	} else {
		fmt.Println("cookie not found from my cart")
	}
}
