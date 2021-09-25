package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName string
	Password string
	First    string
	Last     string
}

var dbUsers = map[string]user{}      //user ID, user
var dbSessions = map[string]string{} //Session ID, user ID

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		// case: username is taken
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "User Name already taken", http.StatusForbidden)
			return
		}
		// Create Session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "my-session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		// Set Link between username and uuid
		dbSessions[c.Value] = un

		//store user in dbUsers
		u := user{un, p, f, l}
		// set Link between user name and user
		dbUsers[un] = u

		//redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return

	}
	// Create Session
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}
