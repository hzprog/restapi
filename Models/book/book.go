package book

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model

	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Image  string `json:"image"`
}
