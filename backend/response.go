package main

import (
	"encoding/json"
	"net/http"

)


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
