package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AndrewBabitt/practice/golang/webservice/controllers"
	"github.com/AndrewBabitt/practice/golang/webservice/models"
)

func main() {
	controllers.RegisterControllers()
	fmt.Println("before for")
	for i := 0; i <= 100; i++ {
		un := fmt.Sprintf("fakeUserName-%v", i)
		em := fmt.Sprintf("fakeEmail-%v@mail.com", i)
		u := models.User{Username: un, Email: em}
		fmt.Println(u)
		models.AddUser(u)
	}
	log.Println("starting server")
	http.ListenAndServe(":3000", nil)
	log.Println("stopping server")
}
