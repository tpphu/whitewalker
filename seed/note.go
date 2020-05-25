package seed

import (
	"time"

	"github.com/icrowley/fake"
	"github.com/tpphu/whitewalker/model"
)

func (s Seeder) NoteSeed() {
	for i := 0; i < 20; i++ {
		n := model.Note{}
		n.ID = 0
		n.Title = fake.FirstName()
		n.CreatedAt = time.Now()
		n.UpdatedAt = time.Now()
		n.DeletedAt = nil
		s.DB.Create(&n)
	}
}
