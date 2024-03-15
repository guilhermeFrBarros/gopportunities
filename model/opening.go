package model

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

// New
func (o *Opening) NewOpening(createDTO *CreateOpeningDTO) {
	o.Role = createDTO.Role
	o.Company = createDTO.Company
	o.Location = createDTO.Location
	o.Remote = *createDTO.Remote
	o.Link = createDTO.Link
	o.Salary = createDTO.Salary
}

func (o *Opening) Updated(updateRequest *UpdateOpeningRequest) {
	if updateRequest.Role != "" {
		o.Role = updateRequest.Role
	}

	if updateRequest.Company != "" {
		o.Company = updateRequest.Company
	}

	if updateRequest.Location != "" {
		o.Location = updateRequest.Location
	}

	if updateRequest.Remote != nil {
		o.Remote = *updateRequest.Remote
	}

	if updateRequest.Link != "" {
		o.Link = updateRequest.Link
	}

	if updateRequest.Salary > 0 {
		o.Salary = updateRequest.Salary
	}
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
