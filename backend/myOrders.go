package main

import (
	"fmt"
	"glow/database"
	"net/http"
)

func GetMyOrderProducts(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session_id")
	if cookie != nil {
		UserID := database.FindUserBySessionId(cookie.Value)
		fmt.Println("User id from myorders - ", UserID)
		OrderedProducts, err := database.DisplayMyOrders(UserID)
		if err != nil {
			WriteError(w, http.StatusBadRequest, err.Error())

		} else {
			fmt.Println("Ordered products to send to frontend  - ", OrderedProducts)
			WriteJson(w, http.StatusAccepted, r, OrderedProducts)
		}
	} else {
		fmt.Println("cookie not found from my orders")
	}
}
