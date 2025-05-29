package main

// import (
// 	"fmt"
// 	"glow/shared"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func productByPid(w http.ResponseWriter, r *http.Request) {
// 	var result = Sha
// 	givenId := mux.Vars(r)
// 	id := givenId["pid"]
// 	fmt.Println("given id", id)
// 	var product shared.Product
// 	var isMatch bool
// 	for _, value := range result {
// 		if value.Pid == id {
// 			isMatch = true
// 			product = value
// 			fmt.Println("Product = ", product)
// 			break
// 		}
// 	}
// 	if !isMatch {
// 		WriteError(w, http.StatusBadRequest, "Invalid id")
// 		return
// 	}
// 	WriteJson(w, http.StatusAccepted, r, product)
// }
