package main

import (
	"module/21_controllers_exercises/02_persist_data/controllers"
	"module/21_controllers_exercises/02_persist_data/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)

}

func getSession() map[string]models.User {
	return models.LoadUsers()
}
