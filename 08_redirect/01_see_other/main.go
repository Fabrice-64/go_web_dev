package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tpl *template.Template

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
	fmt.Print("Your Method as foo: ", req.Method, "\n\n")
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your Method as bar: ", req.Method, "\n\n")
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your Method as barred: ", req.Method, "\n\n")
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
