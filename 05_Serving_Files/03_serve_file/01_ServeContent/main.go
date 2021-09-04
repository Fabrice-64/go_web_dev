package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hedgehog-pic", hedgehogPic)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "Bienvenue sur la page d'accueil")
}

func hedgehogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("herisson.jpeg")
	if err != nil {
		http.Error(w, "file not found", 404)
	}

	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
	}
	http.ServeContent(w, req, fi.Name(), fi.ModTime(), f)

}
