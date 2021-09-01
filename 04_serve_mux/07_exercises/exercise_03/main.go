package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func idx(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is the Index")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is the Dog Page")
}
func me(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("me.gohtml")
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = tpl.ExecuteTemplate(w, "me.gohtml", "Fabrice")
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func main() {

	http.Handle("/", http.HandlerFunc(idx))
	http.Handle("/dog", http.HandlerFunc(dog))
	http.Handle("/me", http.HandlerFunc(me))
	http.ListenAndServe(":8080", nil)

}
