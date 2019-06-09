package handler

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tpphu/whitewalker/mock"
	"github.com/tpphu/whitewalker/model"
)

type UserHanlerTestSuite struct {
	suite.Suite
	userRepo    *mock.UserRepoImpl
	userHandler userHandlerImpl
}

func (suite *UserHanlerTestSuite) SetupTest() {
	userRepo := new(mock.UserRepoImpl)
	suite.userHandler = userHandlerImpl{
		userRepo: userRepo,
	}
	suite.userRepo = userRepo
}

func (suite *UserHanlerTestSuite) TearDownTest() {
}

func TestUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UserHanlerTestSuite))
}

func (suite *UserHanlerTestSuite) TestUserFind() {
	suite.Run("find with valid data", func() {
		var id uint = 49
		out := &model.User{}
		out.ID = id
		out.Name = "Phu"
		suite.userRepo.On("Find", id).Return(out, nil)
		user, _ := suite.userHandler.get(id)
		if user.ID != id {
			suite.Fail("User ID should be same")
		}
	})
	suite.Run("find with invalid data", func() {
		var id uint = 49
		suite.userRepo.On("Find", id).Return(nil, errors.New("Not found"))
		_, err := suite.userHandler.get(id)
		if err != nil {
			suite.Fail("This should be error")
		}
	})
}
