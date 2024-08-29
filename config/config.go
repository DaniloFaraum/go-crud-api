package config

import "gorm.io/gorm"

var (
	db *gorm.DB
	logger *Logger
)

func Init() error { //Starts the Database
	return nil
}

func GetLogger(prefix string) *Logger{
	logger = NewLogger(prefix)
	return logger
}