package book

import (
	"gorm.io/gorm"
)

// Book represents the model for an book
type Book struct {
	gorm.Model

	Isbn   string `json:"isbn" example:"1212154"`
	Title  string `json:"title" example:"title of a book"`
	Author string `json:"author" example:"name of author"`
	Image  string `json:"image" example:"//path to an image"`
}
