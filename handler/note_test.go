package handler

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tpphu/whitewalker/mock"
	"github.com/tpphu/whitewalker/model"
)

type NoteHanlerTestSuite struct {
	suite.Suite
	noteRepo    *mock.NoteRepoImpl
	noteHandler noteHandlerImpl
}

func (suite *NoteHanlerTestSuite) SetupTest() {
	noteRepo := new(mock.NoteRepoImpl)
	suite.noteHandler = noteHandlerImpl{
		noteRepo: noteRepo,
	}
	suite.noteRepo = noteRepo
}

func (suite *NoteHanlerTestSuite) TearDownTest() {
}

func TestNoteRepoTestSuite(t *testing.T) {
	suite.Run(t, new(NoteHanlerTestSuite))
}

func (suite *NoteHanlerTestSuite) TestNoteRepoCreate() {
	suite.Run("create with valid data", func() {
		var id uint = 49
		out := &model.Note{}
		out.ID = id
		suite.noteRepo.On("Find", id).Return(out, nil)
		note, _ := suite.noteHandler.get(id)
		if note.ID != uint(id) {
			suite.Fail("Note ID should be same")
		}
	})
	suite.Run("create with invalid data", func() {
		var id uint = 49
		suite.noteRepo.On("Find", id).Return(nil, errors.New("Not found"))
		_, err := suite.noteHandler.get(id)
		if err != nil {
			suite.Fail("This should be error")
		}
	})
}
