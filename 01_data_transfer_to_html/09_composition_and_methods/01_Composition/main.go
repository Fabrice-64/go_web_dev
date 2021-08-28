package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Spring, Summer, Fall semester
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				{"Info-001", "Introduction to Go", "4"},
				{"Info-002", "Deepen Go", "5"},
				{"Info-003", "Methods in Go", "4"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				{"Info-004", "Deployment with Go", "3"},
				{"Info-005", "Servers with Go", "4"},
				{"Info-006", "Scalability", "4"},
			},
		},
	}
	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}
}
