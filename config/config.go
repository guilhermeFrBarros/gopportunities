package config

import (
	"gorm.io/gorm"
)

var ( // cria uma variavel global
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger(p string) *Logger {
	// Inializer Logger
	logger = NewLogger(p)
	return logger
}
