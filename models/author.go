package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name      string `json:"name"`
	BirthDate string `json:"birthDate"`
}
