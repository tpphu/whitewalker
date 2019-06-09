package handler

import (
	"errors"
	"log"
	"os"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/iris-contrib/httpexpect"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"
	"github.com/stretchr/testify/suite"
)

type RouterTestSuite struct {
	suite.Suite
	SQLMock sqlmock.Sqlmock
	Expect  *httpexpect.Expect
}

func TestRouterTestSuite(t *testing.T) {
	suite.Run(t, new(RouterTestSuite))
}

func (s *RouterTestSuite) SetupTest() {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)

	db, mock, _ := sqlmock.New()
	s.SQLMock = mock
	DB, _ := gorm.Open("mysql", db)

	app := iris.Default()
	initDev(app)
	initNote(app, logger, DB)

	s.Expect = httptest.New(s.T(), app)
}

func (s *RouterTestSuite) TearDownTest() {
}

func (s *RouterTestSuite) TestPing() {
	s.Expect.GET("/ping").Expect().Status(httptest.StatusOK)
}

func (s *RouterTestSuite) TestNote() {
	s.Run("Test find a note", func() {
		var noteID uint = 49
		// Mock du lieu tra ve
		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "title", "completed"}).
			AddRow(noteID, time.Now(), time.Now(), nil, "Todo 123", true)

		// Trong truong ho query
		s.SQLMock.ExpectQuery("SELECT \\* FROM `notes`").
			WillReturnRows(rows)
		expect := s.Expect.GET("/note/" + string(noteID)).Expect()
		expect.Status(httptest.StatusOK)
		expect.JSON().Object().ContainsKey("ID").ValueEqual("ID", 49)
	})
	s.Run("Test find a note", func() {
		var noteID uint = 49
		s.SQLMock.ExpectQuery("SELECT \\* FROM `notes`").
			WillReturnError(errors.New("not found"))
		expect := s.Expect.GET("/note/" + string(noteID)).Expect()
		expect.Status(httptest.StatusNotFound)
	})
}

func (s *RouterTestSuite) TestUser() {
	s.Run("Test find a valid user", func() {
		var userID uint = 50
		// Mock du lieu tra ve
		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name"}).
			AddRow(userID, time.Now(), time.Now(), nil, "Phu")

		// Trong truong ho query
		s.SQLMock.ExpectQuery("SELECT \\* FROM `users`").
			WillReturnRows(rows)
		expect := s.Expect.GET("/user/" + string(userID)).Expect()
		expect.Status(httptest.StatusOK)
		// expect.JSON().Object().ContainsKey("ID").ValueEqual("ID", userID)
	})
	s.Run("Test find an valid user", func() {
		var userID uint = 49
		s.SQLMock.ExpectQuery("SELECT \\* FROM `users`").
			WillReturnError(errors.New("not found"))
		expect := s.Expect.GET("/user/" + string(userID)).Expect()
		expect.Status(httptest.StatusNotFound)
	})
}
