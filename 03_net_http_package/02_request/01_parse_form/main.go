package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

type airplane int

func (a airplane) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

func main() {
	var a airplane
	http.ListenAndServe(":8080", a)
}

// Ne fonctionne pas : pas de retour.
