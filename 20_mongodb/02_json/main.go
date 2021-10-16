package main

// To use in the terminal after launch curl http://localhost:8080/user/1
import (
	"encoding/json"
	"fmt"
	"log"
	"module/20_mongodb/02_json/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	http.ListenAndServe("localhost:8080", r)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
		<html lang="eng">
		<head>
		<meta charset="UTF-8">
		<title>INDEX</title>
		</head>
		<body>
		<h1>Index Page</h1>
		<a href="/user/123456789">GOTO : localhost:8080/user/123456789</a>
		</body>
		</html>`
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func getUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "James Bond",
		Gender: "Male",
		Age:    42,
		Id:     p.ByName("id"),
	}
	uj, err := json.Marshal(u)
	if err != nil {
		log.Fatalln(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}
