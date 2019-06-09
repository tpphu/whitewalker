package repo

import (
	"errors"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
)

type UserRepoTestSuite struct {
	suite.Suite
	userRepo UserRepoImpl
	mock     sqlmock.Sqlmock
}

func (suite *UserRepoTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mock = mock
	userRepo := UserRepoImpl{}
	userRepo.DB, _ = gorm.Open("mysql", db)
	suite.userRepo = userRepo
}

func (suite *UserRepoTestSuite) TearDownTest() {
	suite.userRepo.DB.Close()
}

func TestUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}

func (suite *UserRepoTestSuite) TestUserRepoFind() {
	var userID uint = 5
	suite.Run("find with having found id", func() {
		// Mock du lieu tra ve
		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name"}).
			AddRow(userID, time.Now(), time.Now(), nil, "Phu")

		// Trong truong ho query
		suite.mock.ExpectQuery("SELECT \\* FROM `users`").
			WillReturnRows(rows)

		actual, err := suite.userRepo.Find(userID)
		if err != nil {
			suite.Fail("Error should be nil")
		}
		if actual.ID != userID {
			suite.Fail("Id should be same")
		}
		if actual.DeletedAt != nil {
			suite.Fail("DeletedAt should be nil")
		}
	})

	suite.Run("find with not found id", func() {
		// Trong turong hop khong co cai id
		suite.mock.ExpectQuery("SELECT \\* FROM `users`").
			WillReturnError(errors.New("record not found"))
		_, err := suite.userRepo.Find(userID)
		if err == nil {
			suite.Fail("Error should be not nil")
		}
	})
}
