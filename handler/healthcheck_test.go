package handler

import (
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/iris-contrib/httpexpect"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"
	"github.com/stretchr/testify/suite"
)

type HealthCheckHandlerTestSuite struct {
	suite.Suite
	Expect *httpexpect.Expect
	mock   sqlmock.Sqlmock
	db     *gorm.DB
}

func TestHealthCheckHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HealthCheckHandlerTestSuite))
}

func (s *HealthCheckHandlerTestSuite) SetupTest() {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)

	app := iris.Default()

	db, mock, _ := sqlmock.New()
	DB, _ := gorm.Open("mysql", db)

	s.mock = mock

	noteHanler := &healthCheckHandlerImpl{
		log: logger,
		db:  DB,
	}
	noteHanler.inject(app)
	s.Expect = httptest.New(s.T(), app)
}

func (s *HealthCheckHandlerTestSuite) TearDownTest() {
}

func (s *HealthCheckHandlerTestSuite) TestPing() {
	s.Run("Test ping", func() {
		url := "/ping"
		s.mock.
			ExpectQuery("SELECT 1 as ping").
			WillReturnRows(sqlmock.NewRows([]string{"ping"}).AddRow("1"))

		expect := s.Expect.GET(url).Expect()
		expect.Status(httptest.StatusOK)
		expect.JSON().Object().ContainsMap(map[string]bool{
			"sucess":   true,
			"database": true,
		})
	})
}
