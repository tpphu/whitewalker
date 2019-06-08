package seed

import (
	"github.com/icrowley/fake"
	"github.com/jinzhu/gorm"
	"github.com/tpphu/whitewalker/model"
)

func DepartmentSeed(db *gorm.DB) {
	for i := 0; i < 20; i++ {
		n := model.Department{}
		n.Name = fake.FirstName()
		db.Create(&n)
	}
}
