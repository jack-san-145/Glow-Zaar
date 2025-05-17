package main

import (
	"fmt"
	"net/http"
	"path"
)

func LoadAssert(w http.ResponseWriter, r *http.Request) {
	fmt.Println("working")
	image := path.Base(r.URL.Path)
	im := "./asset/" + image
	http.ServeFile(w, r, im)

}
