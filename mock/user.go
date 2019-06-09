package mock

import (
	"github.com/stretchr/testify/mock"
	"github.com/tpphu/whitewalker/model"
)

type UserRepoImpl struct {
	mock.Mock
}

func (self *UserRepoImpl) Find(id uint) (*model.User, error) {
	args := self.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}
