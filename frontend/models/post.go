package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Post string `json:"post"`
}

type PostResponse struct {
	Data  uint
	Error string
}
