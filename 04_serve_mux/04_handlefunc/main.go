package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Ceci est le dog avec Handlefun")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Ceci est le cat avec Handlefunc")
}

func main() {
	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)
}
