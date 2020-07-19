package main

import (
	"fmt"

	"github.com/AndrewBabitt/practice/golang/webservice/models"
)

func main() {
	u := models.User{ID: 0, Username: "User.Name1", Email: "E@mail.com"}
	fmt.Println(u)
}
