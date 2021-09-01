package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case ("/dog"):
		io.WriteString(w, "Ceci est une première réponse à l'adresse /dog")
	case ("/cat"):
		io.WriteString(w, "Ceci est la réponse à /cat")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
