package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, *req http.Request) {
	io.WriteString(w, "Ici c'est le HandlerFunc de dog")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Ici c'est le HandlerFunc de cat")
}

func main() {
	http.Handle("/dog", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))

	http.ListenAndServe(":8080", nil)
}
