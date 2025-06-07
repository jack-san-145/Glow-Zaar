package main

import (
// 	"context"
// 	"github.com/minio/minio-go/v7"
// 	"github.com/minio/minio-go/v7/pkg/credentials"
// 	"log"
// 	"time"
)

// var (
// 	minioClient *minio.Client
// 	err         error
// 	bucketName  string
// 	expiry      time.Duration
// )

// func minioInitialize() {
// 	// Initialize MinIO client
// 	minioClient, err = minio.New("localhost:9000", &minio.Options{
// 		Creds:  credentials.NewStaticV4("Jack-san", "Jack@145", ""),
// 		Secure: false, // set true if https
// 	})
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	bucketName = "glowzaar-product-images"
// 	// Set expiry time for the signed URL
// 	expiry = time.Minute * 120 // link valid for 2 hrs
// }

// func FindPosterFromPid(product_type *string, poster *string) {

// 	objectName := *product_type + "/" + *poster
// 	presignedURL, err := minioClient.PresignedGetObject(context.Background(), bucketName, objectName, expiry, nil)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	urlStr := presignedURL.String()
// 	poster = &urlStr

// }
