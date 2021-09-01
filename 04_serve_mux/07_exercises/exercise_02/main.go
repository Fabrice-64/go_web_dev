package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func idx(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Page Index")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Page pour le Dog")
}

func me(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("me.gohtml")
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = tpl.ExecuteTemplate(w, "me.gohtml", "Template du Chaton")
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func main() {
	http.HandleFunc("/", idx)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me", me)

	http.ListenAndServe(":8080", nil)
}
