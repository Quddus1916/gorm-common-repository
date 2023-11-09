package stubs

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bondhansarker/gorm-common-repository/mock/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var err error

func NewGormDb() (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}
	dbUrl := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "", "tcp(127.0.0.1:3307)", "gorm")
	opts := &gorm.Config{}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	opts.Logger = newLogger

	db, err = gorm.Open(mysql.Open(dbUrl), opts)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.City{})
	db.AutoMigrate(&models.User{})
}
