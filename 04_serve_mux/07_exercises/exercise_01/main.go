package main

import (
	"io"
	"net/http"
)

func i(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is the Index Page")
}

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is the dog page")
}

func m(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "My Name is Fabrice")
}

func main() {
	http.HandleFunc("/", i)
	http.HandleFunc("/dog", d)
	http.HandleFunc("/me", m)

	http.ListenAndServe(":8080", nil)
}
