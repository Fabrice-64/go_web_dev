package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tpl *template.Template

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your Method at foo: ", req.Method)
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your Method at bar: ", req.Method)
	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your Method at barred: ", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
