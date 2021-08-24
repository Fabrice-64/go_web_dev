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

type car struct {
	Type         string
	Manufacturer string
	Doors        int
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	a := sage{
		Name:  "Buddha",
		Motto: "I'm Wise",
	}
	b := sage{
		Name:  "Jesus",
		Motto: "I'm nice",
	}

	c := car{
		Type:         "coupé",
		Manufacturer: "BMW",
		Doors:        3,
	}

	d := car{
		Type:         "Limousine",
		Manufacturer: "Peugeot",
		Doors:        4,
	}

	sages := []sage{a, b}
	cars := []car{c, d}

	data := struct {
		Wisdom     []sage
		Automobile []car
	}{
		sages,
		cars,
	}

	err := tpl.Execute(os.Stdout, data)

	if err != nil {
		log.Fatalln(err)
	}
}
