package repo

import (
	"github.com/jinzhu/gorm"
	"github.com/tpphu/whitewalker/model"
)

// UserRepo interface
type UserRepo interface {
	Find(uint) (*model.User, error)
}

// UserRepoImpl struct
type UserRepoImpl struct {
	DB *gorm.DB
}

// Find a user
func (userRepo UserRepoImpl) Find(id uint) (*model.User, error) {
	user := &model.User{}
	err := userRepo.DB.Preload("Departments").Find(&user, id).Error
	return user, err
}
