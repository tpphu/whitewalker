package model

import "github.com/jinzhu/gorm"

// Note struct
type Note struct {
	gorm.Model
	Title     string `gorm:"size:32" binding:"required,min=6,max=255"`
	Completed bool
}
