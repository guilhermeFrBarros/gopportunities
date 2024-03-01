package schemas

import (
	"time"

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

type OpeningResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty` //omitempty: omiti o campo  no json se não foi passado nada ou for nil, false, 0
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"Location"`
	Remote    bool      `json:"romote"`
	Link      string    `json:"link"`
	Salary    int64     `json:"salary"`
}
