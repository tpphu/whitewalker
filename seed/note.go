package seed

import (
	"time"

	"../model"
	"github.com/bxcodec/faker"
	"github.com/jinzhu/gorm"
)

func NoteSeed(db *gorm.DB) {
	for i := 0; i < 20; i++ {
		n := model.Note{}
		n.ID = 0
		n.Title = faker.Name()
		n.CreatedAt = time.Now()
		n.UpdatedAt = time.Now()
		n.DeletedAt = nil
		db.Create(&n)
	}
}
