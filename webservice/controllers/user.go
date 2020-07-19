package controllers

import (
	"net/http"
	"regexp"
)

//there are not classic constructors for GO
type userController struct {
	userIDPattern *regexp.Regexp
}

//this allows function to act like a method
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller!"))
}

// this is not a real constructor, it's just a convention to mimic OOP
func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
