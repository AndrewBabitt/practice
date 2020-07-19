package controllers

import (
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()

	// /users different from /users/
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
