package controllers

import (
	"html/template"
	"module/21_controllers_exercises/03_with_templates/session"
	"net/http"
)

type Controller struct {
	tpl *template.Template
}

func NewController(t *template.Template) *Controller {
	return &Controller{t}
}

func (c Controller) Index(w http.ResponseWriter, r *http.Request) {
	u := session.GetUser(w, r)
	session.Show()
	c.tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func (c Controller) Bar(w http.ResponseWriter, r *http.Request) {
	u := session.GetUser(w, r)
	if !session.AlreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter", http.StatusForbidden)
		return
	}
	session.Show() // for Démo
	c.tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
