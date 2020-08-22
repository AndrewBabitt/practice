package main

import (
	"github.com/AndrewBabitt/practice/gorm-example/controllers"
	"github.com/AndrewBabitt/practice/gorm-example/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectToDB()

	user_route := router.Group("/user")
	{
		user_route.GET("/:userId", controllers.GetUser)
		user_route.POST("/new", controllers.CreateUser)
	}
	router.Run()
}
