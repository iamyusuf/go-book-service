package main

import (
	"github.com/gin-gonic/gin"
	"ozone.com/book-service/controllers"
	"ozone.com/book-service/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/authors", controllers.FindAuthors)
	r.POST("/authors", controllers.CreateAuthor)
	r.GET("/authors/:id", controllers.FindAuthor)

	r.Run()
}
