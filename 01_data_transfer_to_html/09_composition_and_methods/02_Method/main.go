package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

type agent struct {
	person
	LicenseToKill bool
}

func (a agent) SomeProcessing() int {
	return 7
}
func (a agent) AgeDbl() int {
	return a.Age * 2
}
func (a agent) TakeArgs(x int) int {
	return x * 2
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	a := agent{
		person{
			Name: "Ian Fleming",
			Age:  56,
		},
		false,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", a)
	if err != nil {
		log.Fatalln(err)
	}

}
