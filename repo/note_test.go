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
	var noteID uint = 5
	suite.Run("create with valid data", func() {
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
	var noteID uint = 5
	suite.Run("find with having found id", func() {
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
	suite.Run("find with not found id", func() {
		suite.mock.ExpectQuery("SELECT \\* FROM `notes`").
			WillReturnError(errors.New("record not found"))
		_, err := suite.noteRepo.Find(int(noteID))
		if err == nil {
			suite.Fail("Error should be not nil")
		}
	})
}

func (suite *NoteRepoTestSuite) TestNoteRepoUpdate() {
	var noteID uint = 5
	suite.Run("update with valid data", func() {
		note := model.Note{
			Title:     "Todo 123",
			Completed: true,
		}
		suite.mock.ExpectExec("UPDATE `notes`").
			WillReturnResult(sqlmock.NewResult(0, 1))
		err := suite.noteRepo.Update(int(noteID), note)
		if err != nil {
			suite.Fail("error should be nil")
		}
	})
	suite.Run("update with invalid data", func() {
		note := model.Note{
			Title:     fake.CharactersN(100),
			Completed: true,
		}
		suite.mock.ExpectExec("UPDATE `notes`").
			WillReturnError(errors.New("Title is exceed 255 character"))
		err := suite.noteRepo.Update(int(noteID), note)
		if err == nil {
			suite.Fail("Error should not nil")
		}
	})
}

// TestNoteRepoDelete used to test delete a note
// it shows that return error has not have meaning
func (suite *NoteRepoTestSuite) TestNoteRepoDelete() {
	var noteID uint = 5
	suite.Run("delete with valid data", func() {
		suite.mock.ExpectExec("UPDATE `notes` SET `deleted_at`=").
			WillReturnResult(sqlmock.NewResult(0, 1))
		err := suite.noteRepo.Delete(int(noteID))
		if err != nil {
			suite.Fail("error should be nil")
		}
	})
	suite.Run("delete with invalid data", func() {
		suite.mock.ExpectExec("UPDATE `notes` SET `deleted_at`=").
			WillReturnResult(sqlmock.NewResult(0, 0))
		err := suite.noteRepo.Delete(int(noteID))
		if err != nil {
			suite.Fail("error should be nil")
		}
	})
}
