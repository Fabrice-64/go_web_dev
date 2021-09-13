package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your Method at foo: ", req.Method)
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your Method at bar; ", req.Method)
	http.Redirect(w, req, "/", http.StatusMovedPermanently)
}
