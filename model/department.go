package model

import "github.com/jinzhu/gorm"

// Department struct
type Department struct {
	gorm.Model
	Name string
}
