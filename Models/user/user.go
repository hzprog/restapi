package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string  `json:"username"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
}
