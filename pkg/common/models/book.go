package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model         // adds ID, created_at etc.
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Description string `json:"description" binding:"required"`
	// User        User
}
