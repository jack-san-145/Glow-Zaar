package main

import (
	"fmt"
	"glow/database"
	"glow/shared"
	"net/http"

	"github.com/gorilla/mux"
)

func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			return // Preflight request
		}

		next(w, r)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/glow-zaar/load-products/{product_type}", withCORS(GetProducts)).Methods("GET")
	// router.HandleFunc("/product-details/{pid}", withCORS(productByPid)).Methods("GET")
	router.HandleFunc("/glow-zaar/GetMyCardProducts", withCORS(GetMyCardProducts)).Methods("GET")
	router.HandleFunc("/glow-zaar/GetMyOrderProducts", withCORS(GetMyOrderProducts)).Methods("GET")
	router.HandleFunc("/glow-zaar/login", withCORS(LoginHandler)).Methods("POST")
	router.HandleFunc("/glow-zaar/register", withCORS(RegistrationHandler)).Methods("POST")
	router.HandleFunc("/glow-zaar/addToCart", withCORS(AddToCart)).Methods("POST")
	router.HandleFunc("/glow-zaar/buyNow", withCORS(BuyNowHandler)).Methods("POST")
	router.HandleFunc("/glow-zaar/home", withCORS(LoadIndex)).Methods("GET")
	router.HandleFunc("/glow-zaar/remove/{pid}", withCORS(RemoveFromCart)).Methods("DELETE")
	router.HandleFunc("/glow-zaar/place-order", withCORS(PlaceOrderFromCart)).Methods("POST")
	router.HandleFunc("/glow-zaar/order-all-cartProducts", withCORS(OrderAllCartProducts)).Methods("GET")
	router.HandleFunc("/glow-zaar/my-profile-status", withCORS(ProfileStatus)).Methods("GET")
	router.HandleFunc("/glow-zaar/my-profile", withCORS(MyProfile)).Methods("GET")
	router.HandleFunc("/glow-zaar/logout", withCORS(Logout)).Methods("GET")
	database.ConnectDB()     // connect to the database
	shared.MinioInitialize() //initialize minIO client

	fmt.Println("server is running")
	serve := http.ListenAndServe(":8989", router)
	if serve != nil {
		fmt.Println("server failure")
	}

}
