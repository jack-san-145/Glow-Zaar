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
	Pid           string `json:"pid"`
	Poster        string `json:"poster"`
	Name          string `json:"name"`
	SKU           string `json:"sku"`
	Price         int    `json:"price"`
	Brand         string `json:"brand"`
	Category      string `json:"category"`
	Color         string `json:"color"`
	Material      string `json:"material"`
	Weight        string `json:"weight"`
	Size          string `json:"size"`
	OriginalPrice string `json:"originalPrice"`
	Sale          bool   `json:"sale"`
	Discount      int    `json:"discount"`
	Ordered       bool   `json:"ordered"`
	AddToCart     bool   `json:"addToCart"`
}

var (
	minioClient *minio.Client
	err         error
	bucketName  string
	expiry      time.Duration
)

var alreadyFetched bool
var result []Product

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
	expiry = time.Minute * 120 // link valid for 2 hrs
}

var Cosmetics = []Product{
	{Pid: "cos-1", Poster: "perfume_1.jpg", Name: "Janan Eau de Parfum", SKU: "JAN-PFM-50ML", Price: 599, Brand: "Janan", Category: "Perfume", Color: "Clear", Material: "Glass bottle / liquid", Weight: "50 ml", Size: "50 ml", OriginalPrice: "₹799", Sale: true, Discount: 25, Ordered: false, AddToCart: false},
	{Pid: "cos-2", Poster: "perfume_2.jpg", Name: "Daisy Eau de Toilette", SKU: "DAI-ET-30ML", Price: 459, Brand: "Daisy", Category: "Perfume", Color: "Pale Yellow", Material: "Glass bottle / liquid", Weight: "30 ml", Size: "30 ml", OriginalPrice: "₹659", Sale: true, Discount: 30, Ordered: false, AddToCart: false},
	{Pid: "cos-3", Poster: "powder_1.jpg", Name: "Pond's Dreamflower Talc", SKU: "PON-TLC-400G", Price: 999, Brand: "Pond's", Category: "Talcum Powder", Color: "White", Material: "Plastic bottle / powder", Weight: "400 g", Size: "400 g", OriginalPrice: "₹1 299", Sale: true, Discount: 23, Ordered: false, AddToCart: false},
	{Pid: "cos-4", Poster: "mascara_1.jpg", Name: "Voluminous Waterproof Mascara", SKU: "VOL-MSC-8ML", Price: 250, Brand: "Voluminous", Category: "Mascara", Color: "Black", Material: "Plastic tube / cream", Weight: "8 ml", Size: "Standard", OriginalPrice: "₹350", Sale: true, Discount: 29, Ordered: false, AddToCart: false},
	{Pid: "cos-5", Poster: "lipstick_1.jpg", Name: "Remee Matte Lipstick-Cherry Red", SKU: "REM-LIP-CR01", Price: 260, Brand: "Remee", Category: "Lipstick", Color: "Cherry Red", Material: "Plastic case / stick", Weight: "3.5 g", Size: "Standard", OriginalPrice: "₹349", Sale: true, Discount: 25, Ordered: false, AddToCart: false},
	{Pid: "cos-6", Poster: "foundation_1.jpg", Name: "Lakmé Invisible Finish Foundation (Shade 04)", SKU: "LAK-FND-25ML-04", Price: 299, Brand: "Lakmé", Category: "Foundation", Color: "Warm Beige", Material: "Glass bottle / liquid", Weight: "25 ml", Size: "25 ml", OriginalPrice: "₹399", Sale: true, Discount: 25, Ordered: false, AddToCart: false},
	{Pid: "cos-7", Poster: "facewash_4.jpg", Name: "Pond's Pure Bright Facewash", SKU: "PON-FW-150ML", Price: 599, Brand: "Pond's", Category: "Face Wash", Color: "Charcoal Grey", Material: "Plastic tube / gel", Weight: "150 ml", Size: "150 ml", OriginalPrice: "₹749", Sale: true, Discount: 20, Ordered: false, AddToCart: false},
	{Pid: "cos-8", Poster: "facewash_3.jpg", Name: "Himalaya Neem Face Wash", SKU: "HIM-FW-100ML", Price: 459, Brand: "Himalaya", Category: "Face Wash", Color: "Green Gel", Material: "Plastic tube / gel", Weight: "100 ml", Size: "100 ml", OriginalPrice: "₹575", Sale: true, Discount: 20, Ordered: false, AddToCart: false},
	{Pid: "cos-9", Poster: "facewash_2.jpg", Name: "Carlton Fresh Foam Cleanser", SKU: "CAR-FW-120ML", Price: 999, Brand: "Carlton", Category: "Face Wash", Color: "White Foam", Material: "Plastic bottle / foam", Weight: "120 ml", Size: "120 ml", OriginalPrice: "₹1 299", Sale: true, Discount: 23, Ordered: false, AddToCart: false},
	{Pid: "cos-10", Poster: "facewash_1.jpg", Name: "Beardo Deep Clean Facewash", SKU: "BRD-FW-200ML", Price: 999, Brand: "Beardo", Category: "Face Wash", Color: "Transparent Gel", Material: "Plastic tube / gel", Weight: "200 ml", Size: "200 ml", OriginalPrice: "₹1 399", Sale: true, Discount: 29, Ordered: false, AddToCart: false},
}

func cosmeticsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("running")
	if !alreadyFetched {
		getImagesfromMinIO()
	}
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

func getImagesfromMinIO() {

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
	alreadyFetched = true
}
