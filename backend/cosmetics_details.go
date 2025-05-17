package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Product struct {
	Pid    int    `json:"pid"`
	Poster string `json:"poster"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
}

var (
	minioClient *minio.Client
	err         error
	bucketName  string
	expiry      time.Duration
)

func minioInitialize() {
	// Initialize MinIO client
	minioClient, err = minio.New("localhost:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("Jack-san", "Jack@145", ""),
		Secure: false, // set true if https
	})
	if err != nil {
		log.Fatalln(err)
	}

	bucketName = "glowzaar-product-images"
	// Set expiry time for the signed URL
	expiry = time.Minute * 120 // link valid for 10 minutes
}

var Cosmetics []Product = []Product{
	{Pid: 1, Poster: "perfume_1.jpg", Name: "Janan", Price: 599},
	{Pid: 2, Poster: "perfume_2.jpg", Name: "Daisy", Price: 459},
	{Pid: 3, Poster: "powder_1.jpg", Name: "Ponds", Price: 999},
	{Pid: 4, Poster: "mascara_1.jpg", Name: "Voluminous Mascara", Price: 250},
	{Pid: 5, Poster: "lipstick_1.jpg", Name: "Remee", Price: 260},
	{Pid: 6, Poster: "foundation_1.jpg", Name: "Lakme Foundation", Price: 299},
	{Pid: 7, Poster: "facewash_4.jpg", Name: "Pond's Facewash", Price: 599},
	{Pid: 8, Poster: "facewash_3.jpg", Name: "Himalaya Neem Face Wash", Price: 459},
	{Pid: 9, Poster: "facewash_2.jpg", Name: "Carlton", Price: 999},
	{Pid: 10, Poster: "facewash_1.jpg", Name: "Beardo", Price: 999},
}

func cosmeticsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("running")
	result := getImagesfromMinIO()
	fmt.Println(result)
	WriteJson(w, http.StatusOK, r, result)
}

func WriteJson(w http.ResponseWriter, status int, r *http.Request, data any) {
	w.Header().Add("Content-Type", "application/json")
	temp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
	w.Write(temp)
}

func getImagesfromMinIO() []Product {

	var result []Product
	for _, product := range Cosmetics {

		objectName := "Cosmetics/" + product.Poster
		presignedURL, err := minioClient.PresignedGetObject(context.Background(), bucketName, objectName, expiry, nil)
		if err != nil {
			log.Fatalln(err)
		}
		urlStr := presignedURL.String()
		product.Poster = urlStr
		// objectName := "products/shoe1.jpg"
		result = append(result, product)
		fmt.Println("Presigned URL:", presignedURL.String())
	}
	return result
}
