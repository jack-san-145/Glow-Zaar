package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"glow/database"
	"glow/shared"
	"log"
	"net/http"
)

var Cached_result = make(map[string][]shared.Product)

func LoadProducts(w http.ResponseWriter, r *http.Request, product_array []shared.Product, product_type string) {
	fmt.Println("Loading products ......")
	var result []shared.Product
	if val, ok := Cached_result[product_type]; ok {
		result = val
		fmt.Println("Cached result ..")
	} else {
		//we don't use here the pointer variable by iterate for loop by index
		for _, product := range product_array {

			objectName := product_type + "/" + product.Poster
			presignedURL, err := minioClient.PresignedGetObject(context.Background(), bucketName, objectName, expiry, nil)
			if err != nil {
				log.Fatalln(err)
			}
			urlStr := presignedURL.String()
			product.Poster = urlStr
			// objectName := "products/shoe1.jpg"
			result = append(result, product)
			// fmt.Println("Presigned URL:", presignedURL.String())
		}
		Cached_result[product_type] = result
	}

	WriteJson(w, http.StatusAccepted, r, result)

}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var (
		product_array []shared.Product
		err           error
	)
	cookie, _ := r.Cookie("session_id")
	if cookie != nil {
		fmt.Println("cookie present in product")
	} else {
		fmt.Println("cookie not present in product")
	}
	fmt.Println("geting products .....")
	value := mux.Vars(r)
	product_type := value["product_type"]

	if val, ok := Cached_result[product_type]; ok {
		fmt.Println("cached Database ..")
		product_array = val
	} else {
		product_array, err = database.GetProductsDb(product_type)
		fmt.Println("Fetched from database....")
		fmt.Println("product_array - ", product_array)
		if err != nil {
			fmt.Println("error from database while fetching data ")
			return
		}
	}
	LoadProducts(w, r, product_array, product_type)

}
