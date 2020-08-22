package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/AndrewBabitt/practice/gorm-example/models"

	"github.com/gin-gonic/gin"
)

//Pattern expected in JSON body
type UserInput struct {
	Username string
	Email    string
}

func GetUser(c *gin.Context) {
	log.Println("finding user by ID", c.Param("userId"))
	var returnUser models.User
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Please use numerical ID to find user"})
	}
	//ORM find user
	returnUser, err = models.GetUserById(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
	}
	c.JSON(http.StatusOK, gin.H{"User": returnUser})
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
