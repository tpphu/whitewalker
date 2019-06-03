package repo

import (
	"errors"
	"testing"
	"time"

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
		var noteID uint = 5
		note := model.Note{
			Title:     "Todo 123",
			Completed: true,
		}
		suite.mock.ExpectExec("INSERT INTO `notes`").
			WillReturnResult(sqlmock.NewResult(
				int64(noteID),
				1,
			))
		actual, err := suite.noteRepo.Create(note)
		if err != nil {
			suite.Fail("error should be nil")
		}
		if actual.ID != noteID {
			suite.Fail("Id should be same")
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
			suite.Fail("Error should not nil")
		}
	})
}

func (suite *NoteRepoTestSuite) TestNoteRepoFind() {
	suite.Run("find with valid data", func() {
		var noteID uint = 5
		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "title", "completed"}).
			AddRow(noteID, time.Now(), time.Now(), nil, "Todo 123", true)
		suite.mock.ExpectQuery("SELECT \\* FROM `notes`").
			WillReturnRows(rows)
		actual, err := suite.noteRepo.Find(int(noteID))
		if err != nil {
			suite.Fail("Error should be nil")
		}
		if actual.ID != noteID {
			suite.Fail("Id should be same")
		}
		if actual.DeletedAt != nil {
			suite.Fail("DeletedAt should be nil")
		}
	})
}
