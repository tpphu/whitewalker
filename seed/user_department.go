package seed

import (
	"github.com/tpphu/whitewalker/model"
)

func (s Seeder) UserDeparmentSeed() {
	user5 := model.User{
		Departments: []model.Department{},
	}
	s.DB.Find(&user5, 5)

	department6 := model.Department{}
	s.DB.Find(&department6, 6)
	user5.Departments = append(user5.Departments, department6)

	department7 := model.Department{}
	s.DB.Find(&department7, 7)
	user5.Departments = append(user5.Departments, department7)

	department8 := model.Department{}
	s.DB.Find(&department8, 8)
	user5.Departments = append(user5.Departments, department8)

	s.DB.Save(&user5)
}
