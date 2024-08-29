package config

import (
	"os"

	"github.com/DaniloFaraum/go-crud-api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath) //Checks if the db already exists
	if os.IsNotExist(err){
		logger.Info("DB file not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)
	}

	//Create and connect with DB
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil{
		logger.Errorf("SQLite could not be initialized: %v", err)
	}
	err = db.AutoMigrate(&schemas.Person{}) //Create or update the 'Person' table based on the 'Person' struct.
	if err != nil{
		logger.Errorf("SQLite AutoMigrate failed: %v", err)
		return nil, err
	}
	return db, nil
}
