package handler

import (
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/stretchr/testify/suite"
	"github.com/tpphu/whitewalker/mock"
	"github.com/tpphu/whitewalker/model"
)

type UserHandlerTestSuite struct {
	suite.Suite
	userRepo *mock.UserRepoImpl
	Expect   *httpexpect.Expect
}

func TestUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

func (s *UserHandlerTestSuite) SetupTest() {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)

	app := iris.Default()
	s.userRepo = new(mock.UserRepoImpl)
	userHanler := &userHandlerImpl{
		userRepo: s.userRepo,
		log:      logger,
	}
	userHanler.inject(app)

	s.Expect = httptest.New(s.T(), app)
}

func (s *UserHandlerTestSuite) TearDownTest() {
}

func (s *UserHandlerTestSuite) TestUser() {
	s.Run("Test find a user has found", func() {
		var userID uint = 49
		out := &model.User{}
		out.ID = userID
		s.userRepo.On("Find", userID).Return(out, nil)
		url := fmt.Sprintf("/user/%d", userID)
		expect := s.Expect.GET(url).Expect()
		expect.Status(httptest.StatusOK)
		expect.JSON().Object().ContainsKey("ID").ValueEqual("ID", userID)
	})
	s.Run("Test find a user not found", func() {
		var userID uint = 50
		out := &model.User{}
		url := fmt.Sprintf("/user/%d", userID)
		s.userRepo.On("Find", userID).Return(out, errors.New("Not found"))
		expect := s.Expect.GET(url).Expect()
		expect.Status(httptest.StatusNotFound)
	})
}
