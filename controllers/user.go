package controllers

import (
	"encoding/json"
	"go-webservice/models"
	"net/http"
	"regexp"
)

type UserController struct {
	userIDPattern *regexp.Regexp
}

func (uc UserController) ServeHTTP (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello from the User Controller!!!!"))
}

func (uc *UserController) getAll(w http.ResponseWriter, r *http.Request)  {
	encondeResponseAsJSON(models.GetUsers(), w)
}

func (uc *UserController) get (id int, w http.ResponseWriter)  {
	u, err := models.GetUserByID(id)
	if err !=nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encondeResponseAsJSON(u, w)
}

func (uc *UserController) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	u, err = models.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encondeResponseAsJSON(u, w)
}

func (uc *UserController) put(id int, w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	if id != u.id {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted user must match ID in URL"))
		return
	}
	encondeResponseAsJSON(u, w)
}

func (uc *UserController) delete(id int, w http.ResponseWriter)  {
	err := models.RemoveUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
}

func (uc *UserController) parseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func newUserController() *UserController {
	return &UserController{
		userIDPattern : regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
