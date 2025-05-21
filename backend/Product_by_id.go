package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func productByPid(w http.ResponseWriter, r *http.Request) {
	givenId := mux.Vars(r)
	id := givenId["pid"]
	fmt.Println("given id", id)
	var product Product
	var isMatch bool
	for _, value := range result {
		if value.Pid == id {
			isMatch = true
			product = value
			fmt.Println("Product = ", product)
			break
		}
	}
	if !isMatch {
		WriteError(w, http.StatusBadRequest, "Invalid id")
		return
	}
	WriteJson(w, http.StatusAccepted, r, product)
}

func WriteError(w http.ResponseWriter, status int, msg string) {
	w.Header().Add("Content-Type", "application/json")
	j, _ := json.Marshal(map[string]string{"Error": msg})
	w.Write(j)

}

func WriteJson(w http.ResponseWriter, status int, r *http.Request, data any) {
	w.Header().Add("Content-Type", "application/json")
	temp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
	w.Write(temp)
}
