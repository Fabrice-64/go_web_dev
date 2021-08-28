package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type item struct {
	Dish, Description string
	Price             float64
}

type meal struct {
	Name  string
	Items []item
}

type menu []meal

type restaurant struct {
	Menus menu
	Name  string
}

type restaurants []restaurant

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	m := restaurants{
		restaurant{
			Name: "Pizza Hut",
			Menus: menu{
				meal{
					Name: "Breakfast",
					Items: []item{
						{Dish: "Frites", Description: "Homemade", Price: 1.2},
					},
				},
				meal{
					Name: "Lunch",
				},
			},
		},

		restaurant{
			Name: "El Tarascon",
			Menus: menu{
				meal{
					Name: "Breakfast",
				},
				meal{
					Name: "Lunch",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, m)

	if err != nil {
		log.Fatalln(err)
	}
}
