package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Ceci est la réponse pour un hotdog")
}

type hotcat int

func (n hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Ceci est la réponse pour le hotcat")
}

func main() {
	var ma hotdog
	var na hotcat

	mux := http.NewServeMux()
	mux.Handle("/dog", ma)
	mux.Handle("/cat", na)

	http.ListenAndServe(":8080", mux)

}
