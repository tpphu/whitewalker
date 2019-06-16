package handler

import (
	"log"
	"os"
	"testing"

	"github.com/iris-contrib/httpexpect"
	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"
	"github.com/stretchr/testify/suite"
)

type HealthCheckHandlerTestSuite struct {
	suite.Suite
	Expect *httpexpect.Expect
}

func TestHealthCheckHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HealthCheckHandlerTestSuite))
}

func (s *HealthCheckHandlerTestSuite) SetupTest() {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)

	app := iris.Default()
	noteHanler := &healthCheckHandlerImpl{
		log: logger,
	}
	noteHanler.inject(app)
	s.Expect = httptest.New(s.T(), app)
}

func (s *HealthCheckHandlerTestSuite) TearDownTest() {
}

func (s *HealthCheckHandlerTestSuite) TestPing() {
	s.Run("Test ping", func() {
		url := "/ping"
		expect := s.Expect.GET(url).Expect()
		expect.Status(httptest.StatusOK)
	})
}
