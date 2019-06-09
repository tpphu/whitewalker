package handler

import (
	"log"
	"os"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/iris-contrib/httpexpect"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"
	"github.com/stretchr/testify/suite"
)

type RouterTestSuite struct {
	suite.Suite
	Expect *httpexpect.Expect
}

func TestRouterTestSuite(t *testing.T) {
	suite.Run(t, new(RouterTestSuite))
}

func (suite *RouterTestSuite) SetupTest() {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)

	db, _, _ := sqlmock.New()
	DB, _ := gorm.Open("mysql", db)

	app := iris.Default()
	initDev(app)
	initNote(app, logger, DB)

	suite.Expect = httptest.New(suite.T(), app)
}

func (suite *RouterTestSuite) TearDownTest() {
}

func (suite *RouterTestSuite) TestPing() {
	suite.Expect.GET("/ping").Expect().Status(httptest.StatusOK)
}
