package config

import (
	"fmt"

	"gorm.io/gorm"
)

var ( // cria uma variavel global
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("error initiaziation sqlite: %v", err)
	}
	return nil
}

func GetSQLite() *gorm.DB {
	return db
}
func GetLogger(p string) *Logger {
	// Inializer Logger
	logger = NewLogger(p)
	return logger
}
