package repo

import (
	"errors"
	"fmt"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/icrowley/fake"
	"github.com/jinzhu/gorm"
	"github.com/tpphu/whitewalker/model"
)

func Test_NoteRepoImpl_Create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	noteRepo := NoteRepoImpl{}
	noteRepo.DB, _ = gorm.Open("mysql", db)
	defer noteRepo.DB.Close()
	var returnID uint = 5
	note := model.Note{
		Title:     "Todo 123",
		Completed: true,
	}
	mock.ExpectExec("INSERT INTO `notes`").WillReturnResult(sqlmock.NewResult(
		int64(returnID),
		1,
	))
	actual, err := noteRepo.Create(note)
	if err != nil {
		t.Fail()
	}
	if actual.ID != returnID {
		t.Fail()
	}
}

func Test_NoteRepoImpl_Create_With_Error(t *testing.T) {
	db, mock, _ := sqlmock.New()
	noteRepo := NoteRepoImpl{}
	noteRepo.DB, _ = gorm.Open("mysql", db)
	defer noteRepo.DB.Close()
	note := model.Note{
		Title:     fake.CharactersN(100),
		Completed: true,
	}
	mock.ExpectExec("INSERT INTO `notes`").WillReturnError(errors.New("Title is too long"))
	actual, err := noteRepo.Create(note)
	fmt.Println("actual:", actual)
	if err == nil {
		t.Fail()
	}
}
