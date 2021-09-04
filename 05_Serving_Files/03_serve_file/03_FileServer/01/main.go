package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/hedgehog", hedgehogPic)

	http.ListenAndServe(":8080", nil)
}

func hedgehogPic(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/herisson.jpeg">`)
}

// Danger: en utilisant 127.0.0.1:8080/ , on obtient le code source
