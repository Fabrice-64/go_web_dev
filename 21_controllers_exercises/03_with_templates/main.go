package main

import (
	"html/template"
	"module/21_controllers_exercises/03_with_templates/controllers"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	c := controllers.NewController(tpl)
	http.HandleFunc("/", c.Index)
	http.HandleFunc("/bar", c.Bar)
	http.HandleFunc("/signup", c.Signup)
	http.HandleFunc("/login", c.Login)
	http.HandleFunc("/logout", c.Logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
