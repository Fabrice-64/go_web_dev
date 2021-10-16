package main

import (
	"encoding/json"
	"fmt"
	"module/20_mongodb/02_json/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Use the terminal with:
// curl http://localhost:8080/user/1
//curl -X POST -H "Content-Type: application/json" -d '{"Name":"James Bond","Gender":"male","Age":32,"Id":"777"}' http://localhost:8080/user
// curl -X DELETE -H "Content-Type: application/json" http://localhost:8080/user/777

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	r.POST("/user", createUser)
	r.DELETE("/user/:id", deleteUser)
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
		<h1>This is the Index Page</h1>
		<a href="/user/123456789">GO TO: http://localhost:8080/user/123456789</a>
		</body>
		</html>`

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
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
	// Marshal into json
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)
}

func createUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// Use the cli to create user
	u := models.User{}
	json.NewDecoder(req.Body).Decode(&u)
	u.Id = "007"
	uj, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) //201
	fmt.Fprintf(w, "%s", uj)
}

func deleteUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	// Use code with cli to handle user.
	fmt.Fprint(w, "Write code to delete user !")
}
