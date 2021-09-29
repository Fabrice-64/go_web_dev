package main

import (
	"net/http"
	"text/template"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

type session struct {
	un           string
	lastActivity time.Time
}

var tpl *template.Template
var dbUsers = map[string]user{}       //user name, user
var dbSessions = map[string]session{} //session ID, session
var dbSessionsCleaned time.Time

const sessionLength = 30

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	dbSessionsCleaned = time.Now()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	showSessions()
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusForbidden)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You are not a 007 agent", http.StatusForbidden)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u user
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		// is there a username ?
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username/password do not match", http.StatusForbidden)
			return
		}
		// Check password
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username/password do not match", http.StatusForbidden)
			return
		}
		// Create Session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "my-session",
			Value: sID.String(),
		}
		c.MaxAge = sessionLength
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSessions() //demo
	tpl.ExecuteTemplate(w, "login.gohtml", u)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	c, _ := req.Cookie("my-session")
	delete(dbSessions, c.Value)
	//remove the cookie
	c = &http.Cookie{
		Name:   "my-session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	//clean up sessions
	if time.Since(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}
	http.Redirect(w, req, "/login", http.StatusSeeOther)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	var u user
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")

		// username taken
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		// create Session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "my-session",
			Value: sID.String(),
		}
		c.MaxAge = sessionLength
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}
		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		u = user{un, bs, f, l, r}
		dbUsers[un] = u
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}
