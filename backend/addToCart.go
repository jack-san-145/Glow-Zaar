package main

import (
	"encoding/json"
	"fmt"
	"glow/database"
	"io"
	"net/http"
)

type ToAddCart struct {
	Pid      string `json:"pid"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

var content ToAddCart

func AddToCart(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session_id")
	if cookie == nil {
		fmt.Println("no cookie found ")
		return
	}
	UserID := database.FindUserBySessionId(cookie.Value)
	body, _ := io.ReadAll(r.Body)
	json.Unmarshal(body, &content)
	fmt.Println("content - ", content)
	database.AddThisToCart(content.Pid, content.Price, content.Quantity, UserID)
	WriteJson(w, http.StatusAccepted, r, "successfully inserted to cart ")
}
