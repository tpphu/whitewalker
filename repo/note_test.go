package repo

import (
	"errors"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/icrowley/fake"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
	"github.com/tpphu/whitewalker/model"
)

type NoteRepoTestSuite struct {
	suite.Suite
	noteRepo NoteRepoImpl
	mock     sqlmock.Sqlmock
}

func (suite *NoteRepoTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mock = mock
	noteRepo := NoteRepoImpl{}
	noteRepo.DB, _ = gorm.Open("mysql", db)
	suite.noteRepo = noteRepo
}

func (suite *NoteRepoTestSuite) TearDownTest() {
	suite.noteRepo.DB.Close()
}

func TestNoteRepoTestSuite(t *testing.T) {
	suite.Run(t, new(NoteRepoTestSuite))
}

func (suite *NoteRepoTestSuite) TestNoteRepoCreate() {
	suite.Run("create with valid data", func() {
		var returnID uint = 5
		note := model.Note{
			Title:     "Todo 123",
			Completed: true,
		}
		suite.mock.ExpectExec("INSERT INTO `notes`").
			WillReturnResult(sqlmock.NewResult(
				int64(returnID),
				1,
			))
		actual, err := suite.noteRepo.Create(note)
		if err != nil {
			suite.T().Fail()
		}
		if actual.ID != returnID {
			suite.T().Fail()
		}
	})
	suite.Run("create with invalid data", func() {
		note := model.Note{
			Title:     fake.CharactersN(100),
			Completed: true,
		}
		suite.mock.ExpectExec("INSERT INTO `notes`").
			WillReturnError(errors.New("Title is exceed 255 character"))
		_, err := suite.noteRepo.Create(note)
		if err == nil {
			suite.T().Fail()
		}
	})
}
