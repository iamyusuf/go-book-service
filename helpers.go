package main

import (
	"github.com/gin-gonic/gin"
	"ozone.com/book-service/controllers"
)

func DefineRoutes(r *gin.Engine) {
	// Define and set up routes here
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/authors", controllers.FindAuthors)
	r.POST("/authors", controllers.CreateAuthor)
	r.GET("/authors/:id", controllers.FindAuthor)
}
