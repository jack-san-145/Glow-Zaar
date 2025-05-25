package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			return // Preflight request
		}

		next(w, r)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/load-products/{product_type}", withCORS(getProducts)).Methods("GET")
	router.HandleFunc("/product-details/{pid}", withCORS(productByPid)).Methods("GET")
	minioInitialize() //initialize minIO client
	fmt.Println("server is running")
	serve := http.ListenAndServe(":8989", router)
	if serve != nil {
		fmt.Println("server failure")
	}

}
