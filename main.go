package main

import (
	"project/rest_api/controller"
	"project/rest_api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)

	r.Run()
}
