package controllers

import (
	"net/http"
	"regexp"
)

type UserController struct {
	userIDPatter *regexp.Regexp
}

func (uc UserController) ServeHTTP (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello from the User Controller!!!!"))
}

func newUserController() *UserController {
	return &UserController{
		userIDPatter : regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
