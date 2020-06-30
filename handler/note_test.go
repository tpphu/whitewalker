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

type NoteHandlerTestSuite struct {
	suite.Suite
	noteRepo *mock.NoteRepoImpl
	Expect   *httpexpect.Expect
}

func TestNoteHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(NoteHandlerTestSuite))
}

func (s *NoteHandlerTestSuite) SetupTest() {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)

	app := iris.Default()
	s.noteRepo = new(mock.NoteRepoImpl)
	noteHanler := &noteHandlerImpl{
		noteRepo: s.noteRepo,
		log:      logger,
	}
	noteHanler.inject(app)
	s.Expect = httptest.New(s.T(), app)
}

func (s *NoteHandlerTestSuite) TearDownTest() {
}

func (s *NoteHandlerTestSuite) TestNote() {
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
