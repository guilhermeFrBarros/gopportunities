package model

import (
	"fmt"

	svc "github.com/guilhermeFrBarros/gopportunities/services"
)

type CreateOpeningDTO struct {
	Role     string `json:"role"`
	Company  string `json: "company"`
	Location string `json: "location"`
	Remote   *bool  `json: "remote"`
	Link     string `json: "link"`
	Salary   int64  `json: "salary"`
}

// indica que o um func de CreateOpeningDTO
func (r *CreateOpeningDTO) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" &&
		r.Link == "" && r.Remote == nil && r.Salary >= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Role == "" {
		return svc.ErrParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return svc.ErrParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return svc.ErrParamIsRequired("location", "string")
	}
	if r.Link == "" {
		return svc.ErrParamIsRequired("link", "string")
	}
	if r.Remote == nil {
		return svc.ErrParamIsRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return svc.ErrParamIsRequired("salary", "int64")
	}

	return nil
}
