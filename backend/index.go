package main

import (
	"context"
	"fmt"
	"glow/database"
	"glow/shared"
	"log"
	"net/http"
)

var UserID int
var result_Category []shared.Category

var totatProductType []shared.Category

func loadCategory(w http.ResponseWriter, r *http.Request, totalProductType *[]shared.Category) {
	if len(result_Category) == 0 {
		for _, HomeCard := range *totalProductType {
			objectName := HomeCard.Poster
			presignedURL, err := minioClient.PresignedGetObject(context.Background(), bucketName, objectName, expiry, nil)
			if err != nil {
				log.Fatalln("error while fetching category card from minio")
			}
			urlStr := presignedURL.String()
			HomeCard.Poster = urlStr
			result_Category = append(result_Category, HomeCard)
		}
		fmt.Println("from minio")
	}

	WriteJson(w, http.StatusAccepted, r, result_Category)
}

func LoadIndex(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session_id")
	if cookie != nil {
		UserID = database.FindUserBySessionId(cookie.Value)
		fmt.Println("UserID - ", UserID)
	} else {
		fmt.Println("User id not found")
	}
	fmt.Println("comming into the home page")
	if len(totatProductType) == 0 {
		totatProductType = database.GetAllCategory()
		fmt.Println("totatProductType - ", totatProductType)
	}

	loadCategory(w, r, &totatProductType)
}

func insCart() {
	category := []shared.Category{
		{Product_type_id: "cosmetics", Poster: "cosmetics.jpg", IsProduct: false},
		{Product_type_id: "jewels", Poster: "jewels.jpg", IsProduct: false},
	}
	database.InsertHomeProduct(category)
}
