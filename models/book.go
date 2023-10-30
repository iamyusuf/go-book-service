package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title   string    `json:"title"`
	Authors []*Author `gorm:"many2many:book_authors;"`
}

type Author struct {
	gorm.Model
	Name  string  `json:"name"`
	Books []*Book `gorm:"many2many:book_authors;"`
}

type BookAuthor struct {
	AuthorId uint `json:"authorId"`
	BookId   uint `json:"bookId"`
}

func NewBook(title string, author string) *Book {
	return &Book{
		Title: title,
	}
}

func NewAuthor(name string) *Author {
	return &Author{
		Name: name,
	}
}
