package session

import (
	"fmt"
	"module/21_controllers_exercises/03_with_templates/models"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

const Length int = 30

var Users = map[string]models.User{}       // user ID, user
var Sessions = map[string]models.Session{} //session ID, session with user ID
var LastCleaned time.Time = time.Now()

func GetUser(w http.ResponseWriter, r *http.Request) models.User {
	// get cookie
	c, err := r.Cookie("my-session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "my-session",
			Value: sID.String(),
		}
		c.MaxAge = Length
		http.SetCookie(w, c)
	}
	// If Cookie already exists: get User
	var u models.User
	if s, ok := Sessions[c.Value]; ok {
		s.LastActivity = time.Now()
		Sessions[c.Value] = s
		u = Users[s.UserName]
	}
	return u
}

func AlreadyLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	c, err := r.Cookie("my-session")
	if err != nil {
		return false
	}
	s, ok := Sessions[c.Value]
	if ok {
		s.LastActivity = time.Now()
		Sessions[c.Value] = s
	}
	_, ok = Users[s.UserName]
	// Refresh Session
	c.MaxAge = Length
	http.SetCookie(w, c)
	return ok
}

func Clean() {
	fmt.Println("BEFORE CLEAN") // for Demonstration
	Show()                      // for demonstration
	for k, v := range Sessions {
		if time.Now().Sub(v.LastActivity) > (time.Second * 30) {
			delete(Sessions, k)
		}
		LastCleaned = time.Now()
		fmt.Println("AFTER CLEAN")
		Show()
	}
}

func Show() {
	fmt.Println("**********")
	for k, v := range Sessions {
		fmt.Println(k, v.UserName)
	}
	fmt.Println("")
}
