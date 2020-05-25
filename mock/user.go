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
	// Todo: args.Get(0) can be nil
	return args.Get(0).(*model.User), args.Error(1)
}
