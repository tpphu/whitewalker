package handler

import (
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/iris-contrib/httpexpect"
	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"
	"github.com/stretchr/testify/suite"
	"github.com/tpphu/whitewalker/mock"
	"github.com/tpphu/whitewalker/model"
)

type RouterTestSuite struct {
	suite.Suite
	noteRepo *mock.NoteRepoImpl
	userRepo *mock.UserRepoImpl
	Expect   *httpexpect.Expect
}

func TestRouterTestSuite(t *testing.T) {
	suite.Run(t, new(RouterTestSuite))
}

func (s *RouterTestSuite) SetupTest() {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)

	app := iris.Default()
	initDev(app)
	s.noteRepo = new(mock.NoteRepoImpl)
	noteHanler := &noteHandlerImpl{
		noteRepo: s.noteRepo,
		log:      logger,
	}
	initNote(app, noteHanler)
	s.userRepo = new(mock.UserRepoImpl)
	userHanler := &userHandlerImpl{
		userRepo: s.userRepo,
		log:      logger,
	}
	initUser(app, userHanler)

	s.Expect = httptest.New(s.T(), app)
}

func (s *RouterTestSuite) TearDownTest() {
}

func (s *RouterTestSuite) TestPing() {
	s.Expect.GET("/ping").Expect().Status(httptest.StatusOK)
}

func (s *RouterTestSuite) TestNote() {
	s.Run("Test find a note has found", func() {
		var noteID uint = 49
		out := &model.Note{}
		out.ID = noteID
		s.noteRepo.On("Find", noteID).Return(out, nil)
		url := fmt.Sprintf("/note/%d", noteID)
		expect := s.Expect.GET(url).Expect()
		expect.Status(httptest.StatusOK)
		expect.JSON().Object().ContainsKey("ID").ValueEqual("ID", noteID)
	})
	s.Run("Test find a note not found", func() {
		var noteID uint = 50
		out := &model.Note{}
		s.noteRepo.On("Find", noteID).Return(out, errors.New("Not found"))
		url := fmt.Sprintf("/note/%d", noteID)
		expect := s.Expect.GET(url).Expect()
		expect.Status(httptest.StatusNotFound)
	})
}

func (s *RouterTestSuite) TestUser() {
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
