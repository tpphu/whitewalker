package seed

import (
	"github.com/icrowley/fake"
	"github.com/tpphu/whitewalker/model"
)

func (s Seeder) DepartmentSeed() {
	for i := 0; i < 20; i++ {
		n := model.Department{}
		n.Name = fake.FirstName()
		s.DB.Create(&n)
	}
}
