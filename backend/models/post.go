package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Post  string `json:"post"`
}