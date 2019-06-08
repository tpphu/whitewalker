package seed

import (
	"github.com/jinzhu/gorm"
	"github.com/tpphu/whitewalker/model"
)

func UserDeparmentSeed(db *gorm.DB) {
	user5 := model.User{
		Departments: []model.Department{},
	}
	db.Find(&user5, 5)

	department6 := model.Department{}
	db.Find(&department6, 6)
	user5.Departments = append(user5.Departments, department6)

	department7 := model.Department{}
	db.Find(&department7, 7)
	user5.Departments = append(user5.Departments, department7)

	department8 := model.Department{}
	db.Find(&department8, 8)
	user5.Departments = append(user5.Departments, department8)

	db.Save(&user5)
}
