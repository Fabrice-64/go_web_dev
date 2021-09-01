package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Ceci est le Hotdog avec Handle")
}

type hotcat int

func (n hotcat) ServeHTTP(w http.ResponseWriter, requ *http.Request) {
	io.WriteString(w, "Ceci est la réponse Hotcat avec Handle")
}

func main() {
	var ma hotdog
	var na hotcat

	http.Handle("/dog", ma)
	http.Handle("/cat", na)

	http.ListenAndServe(":8080", nil)
}
