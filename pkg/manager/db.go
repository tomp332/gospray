package manager

import (
	"fmt"
	"github.com/tomp332/gobrute/pkg/internalTypes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var MainDB *gorm.DB

func InitManagerDB() {
	var err error
	MainDB, err = gorm.Open(sqlite.Open(GlobalSettings.DBFilePath), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %s", err.Error()))
	}
	err = MainDB.AutoMigrate(&internalTypes.SlaveDTO{})
	if err != nil {
		panic(fmt.Sprintf("Unable to migrate Slaves table: %s", err.Error()))
	}
	err = MainDB.AutoMigrate(&internalTypes.CredentialsDTO{})
	if err != nil {
		panic(fmt.Sprintf("Unable to migrate Credentials table: %s", err.Error()))
	}
	err = MainDB.AutoMigrate(&internalTypes.BruteForceDTO{})
	if err != nil {
		panic(fmt.Sprintf("Unable to migrate Brute Force table: %s", err.Error()))
	}
	fmt.Println("Successfully connected to database")
}
