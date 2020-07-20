package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()

	// /users different from /users/
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
