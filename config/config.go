package config

import "gorm.io/gorm"

var (
	db *gorm.DB
)

func Init() error { //Starts the Database
	return nil
}
