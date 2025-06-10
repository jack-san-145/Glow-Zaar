package main

import (
	"encoding/json"
	"fmt"
	"glow/database"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type ToPlaceOrder struct {
	Pid      string `json:"pid"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
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

func RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session_id")
	if cookie != nil {
		data := mux.Vars(r)
		pid := data["pid"]
		fmt.Println("pid from remove - ", pid)
		UserID := database.FindUserBySessionId(cookie.Value)
		fmt.Println("User id from removal of mycart - ", UserID)
		database.RemoveFromCart(pid, UserID)
		WriteJson(w, http.StatusAccepted, r, "product removed")
	}
}

func PlaceOrderFromCart(w http.ResponseWriter, r *http.Request) {
	var content ToPlaceOrder
	cookie, _ := r.Cookie("session_id")
	if cookie != nil {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("error while reading data from place order")
			return
		}
		json.Unmarshal(body, &content)

		fmt.Println("content from place-order - ", content)
		UserID := database.FindUserBySessionId(cookie.Value)
		fmt.Println("User id from place-order of mycart - ", UserID)
		database.PlaceOrderFromCartDb(content.Pid, content.Price, content.Quantity, UserID)
		WriteJson(w, http.StatusAccepted, r, " Your order has been Placed")
	}
}

func OrderAllCartProducts(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session_id")
	if cookie != nil {

		UserID := database.FindUserBySessionId(cookie.Value)
		fmt.Println("User id from order all cart products - ", UserID)
		database.OrderAllCartProductsDb(UserID)
		WriteJson(w, http.StatusAccepted, r, "All the cart products has been ordered ")
	}
}
