package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/aboundance", aboundance)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "decided value",
		Path:  "/",
	})
	fmt.Fprintln(w, "CHECK YOUR BROWSER INSPECTOR")
}

func aboundance(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "1st Cookie",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "2nd Cookie",
	})

	fmt.Fprintln(w, "CHECK YOUR BROWSER INSPECTOR")
}

func read(w http.ResponseWriter, req *http.Request) {
	c1, err := req.Cookie("my-cookie")
	if err != nil {
		log.Print(err)
	} else {
		fmt.Fprintln(w, "COOKIE #1: ", c1)
	}
	c2, err := req.Cookie("general")
	if err != nil {
		log.Print(err)
	} else {
		fmt.Fprintln(w, "COOKIE #2: ", c2)
	}

	c3, err := req.Cookie("specific")
	if err != nil {
		log.Print(err)
	} else {
		fmt.Fprintln(w, "COOKIE #3: ", c3)
	}
}
