package controllers

import (
	"module/21_controllers_exercises/03_with_templates/models"
	"module/21_controllers_exercises/03_with_templates/session"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func (c Controller) Signup(w http.ResponseWriter, req *http.Request) {
	if session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	var u models.User
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")
		// Username taken ?
		if _, ok := session.Users[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		//create Session
		sID := uuid.NewV4()
		ck := &http.Cookie{
			Name:  "my-session",
			Value: sID.String(),
		}
		ck.MaxAge = session.Length
		http.SetCookie(w, ck)
		session.Sessions[ck.Value] = models.Session{
			UserName:     un,
			LastActivity: time.Now()}
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Status Server Internal Error", http.StatusInternalServerError)
		}
		u = models.User{
			UserName: un,
			Password: bs,
			First:    f,
			Last:     l,
			Role:     r}
		session.Users[un] = u
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	session.Show()
	c.tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func (c Controller) Login(w http.ResponseWriter, req *http.Request) {
	if session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u models.User
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		// Check if username
		u, ok := session.Users[un]
		if !ok {
			http.Error(w, "This Username does not exist", http.StatusForbidden)
			return
		}
		// Check the password
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Bad Password", http.StatusForbidden)
			return
		}
		sID := uuid.NewV4()
		ck := &http.Cookie{
			Name:  "my-session",
			Value: sID.String(),
		}
		ck.MaxAge = session.Length
		http.SetCookie(w, ck)
		session.Sessions[ck.Value] = models.Session{
			UserName:     un,
			LastActivity: time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	session.Show()
	c.tpl.ExecuteTemplate(w, "login.gohtml", u)
}

func (c Controller) Logout(w http.ResponseWriter, req *http.Request) {
	if !session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	ck, _ := req.Cookie("my-session")
	// delete the session
	delete(session.Sessions, ck.Value)
	// remove the cookie
	ck = &http.Cookie{
		Name:   "my-session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, ck)
	// Clean up Sessions
	if time.Now().Sub(session.LastCleaned) > (time.Second * 30) {
		go session.Clean()
	}
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
