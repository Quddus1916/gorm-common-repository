package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	gorm_common_repository "github.com/bondhansarker/gorm-common-repository"
	"github.com/bondhansarker/gorm-common-repository/mock/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGet(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      mockDb,
		SkipInitializeWithVersion: true,
	})
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Errorf("Failed to open connection to DB: %v", err)
	}

	if db == nil {
		t.Error("Failed to open connection to DB: conn is nil")
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users`").
		WithArgs("nafi", "dhaka", 1).
		WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()
	t.Run("Create", func(t *testing.T) {

		userRequest := models.User{
			Id:   1,
			Name: "nafi",
			City: "dhaka",
		}

		UserRepo := gorm_common_repository.NewCommonRepository[models.User]("users", db)

		_, err := UserRepo.CreateRecord(userRequest)
		if err != nil {
			t.Errorf("error was not expected while creating user: %s", err)
		}

		err = mock.ExpectationsWereMet()
		assert.Nil(t, err)
		t.Logf("test successfull with Error:%s", err)

	})
}
