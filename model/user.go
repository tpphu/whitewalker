package model

import "github.com/jinzhu/gorm"

// User struct
type User struct {
	gorm.Model
	Name        string
	Departments []Department `gorm:"many2many:user_departments;"`
}
