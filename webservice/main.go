package main

import (
	"net/http"

	"github.com/AndrewBabitt/practice/golang/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
