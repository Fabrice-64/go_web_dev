package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name   string
	City   string
	Zip    string
	Street string
}

type region struct {
	Name   string
	Hotels []hotel
}

type Regions []region

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	h := Regions{
		region{
			Name: "Southern",
			Hotels: []hotel{
				{Name: "Ibis", City: "Marseille", Street: "Canebière", Zip: "13000"},
				{Name: "Mercure", City: "Marignane", Street: "Aéroport", Zip: "13000"},
			},
		},

		region{
			Name: "Central",
			Hotels: []hotel{
				{Name: "Novotel", City: "Orléans", Street: "UNK", Zip: "45000"},
				{Name: "Sofitel", City: "Paris", Street: "UNK", Zip: "75000"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
