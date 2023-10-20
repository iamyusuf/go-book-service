package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ozone.com/book-service/models"
)

type CreateAuthorInput struct {
	Name string `json:"name" binding:"required"`
}

func FindAuthors(c *gin.Context) {
	var authors []models.Author
	models.DB.Find(&authors)

	c.JSON(http.StatusOK, gin.H{"authors": authors})
}

func FindAuthor(c *gin.Context) {
	authorId := c.Param("id")
	var author models.Author
	models.DB.Find(&author, authorId)

	c.JSON(http.StatusOK, gin.H{"author": author})
}

func CreateAuthor(c *gin.Context) {
	// Validate input
	var input CreateAuthorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	author := models.Author{Name: input.Name}
	models.DB.Create(&author)

	c.JSON(http.StatusOK, gin.H{"data": author})
}
