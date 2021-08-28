package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "I love peace",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Follow me",
	}

	john := sage{
		Name:  "John",
		Motto: "Life is life",
	}

	sages := []sage{buddha, jesus, john}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)

	if err != nil {
		log.Fatalln(err)
	}
}
