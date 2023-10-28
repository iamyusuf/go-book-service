package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
}

func NewBook(title string, author string) *Book {
	return &Book{
		Title:  title,
		Author: author,
	}
}
