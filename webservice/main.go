package main

import (
	"log"
	"net/http"

	"github.com/AndrewBabitt/practice/golang/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	log.Println("starting server")
	http.ListenAndServe(":3000", nil)
	log.Println("stopping server")
}
