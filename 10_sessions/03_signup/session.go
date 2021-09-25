package main

import "net/http"

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("my-session")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}

func getUser(req *http.Request) user {
	var u user
	// get cookie
	c, err := req.Cookie("my-session")
	if err != nil {
		return u
	}
	// if user already exists, get user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}
