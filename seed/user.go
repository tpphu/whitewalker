package seed

import (
	"github.com/icrowley/fake"
	"github.com/jinzhu/gorm"
	"github.com/tpphu/whitewalker/model"
)

func UserSeed(db *gorm.DB) {
	for i := 0; i < 20; i++ {
		n := model.User{}
		n.Name = fake.FirstName()
		db.Create(&n)
	}
}
