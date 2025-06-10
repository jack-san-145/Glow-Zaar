package main

import (
	"encoding/json"
	"fmt"
	"glow/database"
	"io"
	"net/http"
)

type BuyNow struct {
	Pid      string `json:"pid"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

func BuyNowHandler(w http.ResponseWriter, r *http.Request) {
	var content BuyNow
	cookie, _ := r.Cookie("session_id")
	if cookie == nil {
		fmt.Println("no cookie found ")
		return
	}
	UserID := database.FindUserBySessionId(cookie.Value)
	body, _ := io.ReadAll(r.Body)
	json.Unmarshal(body, &content)
	fmt.Println("content from buy now - ", content)
	content.Price = content.Price * content.Quantity
	database.BuyItNowDb(content.Pid, content.Quantity, content.Price, UserID)
	WriteJson(w, http.StatusAccepted, r, "successfully inserted to MyOrders ")
}
