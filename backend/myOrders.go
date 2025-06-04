package main

import (
	"fmt"
	"net/http"
)

func getMyOrders(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session_id")
	if cookie != nil {
		fmt.Println("cookie present on orders - ", cookie)
	} else {
		fmt.Println("cookie not present on orders !! ")
	}
}
