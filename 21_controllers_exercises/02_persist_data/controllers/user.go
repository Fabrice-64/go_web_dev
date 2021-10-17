package controllers

import (
	"encoding/json"
	"fmt"
	"module/21_controllers_exercises/02_persist_data/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
)

type UserController struct {
	session map[string]models.User
}

func NewUserController(m map[string]models.User) *UserController {
	return &UserController{m}
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)
	u.Id = uuid.NewV4().String()
	uc.session[u.Id] = u
	models.StoreUser(uc.session)
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", uj)
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	Id := p.ByName("id")
	u := uc.session[Id]
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	Id := p.ByName("id")
	delete(uc.session, Id)
	models.StoreUser(uc.session)

	w.WriteHeader(http.StatusOK)
	fmt.Println("Deleted User: ", Id)
}
