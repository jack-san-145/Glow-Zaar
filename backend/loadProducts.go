package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func loadProducts(w http.ResponseWriter, r *http.Request, product_array []Product, product_type string) {
	fmt.Println("Loading products ......")
	var result []Product

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
	// alreadyFetched = true
	WriteJson(w, http.StatusAccepted, r, result)

}

func getProducts(w http.ResponseWriter, r *http.Request) {
	var product_array []Product
	fmt.Println("geting products .....")
	value := mux.Vars(r)
	product_type := value["product_type"]
	if product_type == "Cosmetics" {
		product_array = Cosmetics
	} else if product_type == "jewels" {
		product_array = Jewels
	}
	loadProducts(w, r, product_array, product_type)

}
