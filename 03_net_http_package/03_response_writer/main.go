package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Fabrice-Key", "This was made by Fabrice")
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	fmt.Fprintln(w, "<h1>This is a plain text sentence</h1>")
}

func main() {
	var t hotdog
	http.ListenAndServe(":8080", t)
}

// localhost:8080 puis aller dans la console pour voir les headers
