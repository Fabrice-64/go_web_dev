package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	sages := map[string]string{
		"Asia":        "Buddha",
		"Middle East": "Jesus",
		"Europe":      "Descartes",
		"Africa":      "Saint Augustin",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)

	if err != nil {
		log.Fatalln(err)
	}
}
