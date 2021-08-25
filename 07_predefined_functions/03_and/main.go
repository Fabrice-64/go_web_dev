package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	Name  string
	Motto string
	Admin bool
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	user1 := user{
		Name:  "Alphonse",
		Motto: "Toujours Prêt",
		Admin: true,
	}

	user2 := user{
		Name:  "François",
		Motto: "Peu mais bien",
		Admin: false,
	}

	user3 := user{
		Name:  "",
		Motto: "Au Taquet",
		Admin: true,
	}

	users := []user{user1, user2, user3}

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}
