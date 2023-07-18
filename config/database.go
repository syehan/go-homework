// @/config/database.go
package config

import (
	"go-homework/models"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectSqlite() error {
	var err error
	var SqliteDb string = os.Getenv("SQLITE_DB")

	Database, err = gorm.Open(sqlite.Open(SqliteDb), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&models.Author{})
	Database.AutoMigrate(&models.Book{})
	Database.AutoMigrate(&models.User{})

	return nil
}
