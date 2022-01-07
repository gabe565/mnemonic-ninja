package internal

import (
	"github.com/gabe565/mnemonic-ninja/internal/word"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func SetupDatabase() (*gorm.DB, error) {
	var err error

	l := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: l,
	})
	if err != nil {
		return db, err
	}

	err = db.AutoMigrate(&word.WordModel{})
	if err != nil {
		return db, err
	}

	return db, nil
}
