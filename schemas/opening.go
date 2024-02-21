package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model // é um atalho. Implicitamente insere os campos ID, CreatedAt, UpdatedAt, DeletedAt
	Role       string
	Company    string
	Location   string
	Remote     bool
	Link       string
	Salary     int64
}
