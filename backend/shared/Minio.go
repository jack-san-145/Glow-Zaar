package shared

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"time"
)

var (
	MinioClient *minio.Client
	Err         error
	BucketName  string
	Expiry      time.Duration
)

func MinioInitialize() {
	// Initialize MinIO client  192.168.250.106
	MinioClient, Err = minio.New("192.168.250.106:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("Jack-san", "Jack@145", ""),
		Secure: false, // set true if https
	})
	if Err != nil {
		log.Fatalln(Err)
	}

	BucketName = "glowzaar-product-images"
	// Set expiry time for the signed URL
	Expiry = time.Minute * 120 // link valid for 2 hrs
}

func FindPosterFromPid(product_type string, poster string) string {

	objectName := product_type + "/" + poster
	presignedURL, err := MinioClient.PresignedGetObject(context.Background(), BucketName, objectName, Expiry, nil)
	if err != nil {
		log.Fatalln(err)
	}
	urlStr := presignedURL.String()
	poster = urlStr
	return poster

}
