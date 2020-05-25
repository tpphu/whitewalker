package seed

import (
	"github.com/icrowley/fake"
	"github.com/tpphu/whitewalker/model"
)

func (s Seeder) UserSeed() {
	for i := 0; i < 20; i++ {
		n := model.User{}
		n.Name = fake.FirstName()
		s.DB.Create(&n)
	}
}
