package controllers

import (
	"net/http"

	"github.com/AndrewBabitt/practice/gorm-example/models"

	"github.com/gin-gonic/gin"
)

//Pattern expected in JSON body
type UserInput struct {
	Username string
	Email    string
}

func GetUser(c *gin.Context) {
	//ORM find user
	models.GetUserById(1)
	c.JSON(http.StatusOK, gin.H{"Status": "OK"})
}

func CreateUser(c *gin.Context) {
	var input UserInput

	//try to assign body to UserInput struct
	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// assign fields to User Model
	user := models.User{Username: input.Username, Email: input.Email}

	// create user
	models.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{"Status": "OK"})
}
