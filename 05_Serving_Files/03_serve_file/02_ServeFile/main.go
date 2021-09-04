package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hedgehog-pic", hedgehogPic)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "Bienvenue sur la page d'accueil ServeFile")
}

func hedgehogPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "herisson.jpeg")
}
