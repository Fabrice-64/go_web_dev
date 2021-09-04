package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hedgehog", hedgehog)
	http.HandleFunc("/hedgehog-pic", hedgehogPic)
	http.ListenAndServe(":9090", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "Bienvenue sur mon mini site")
}

func hedgehog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<!--Internally not served-->
	<img src="herisson.jpeg">
	`)
}

func hedgehogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("herisson.jpeg")
	if err != nil {
		http.Error(w, "Error File not Found - Sorry", 404)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}
