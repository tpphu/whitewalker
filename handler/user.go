package handler

import (
	"github.com/tpphu/whitewalker/model"
	"github.com/tpphu/whitewalker/repo"
)

type userHandlerImpl struct {
	userRepo repo.UserRepo
}

func (n userHandlerImpl) get(id uint) (*model.User, error) {
	return n.userRepo.Find(id)
}
