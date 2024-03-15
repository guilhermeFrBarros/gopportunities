package model

import "fmt"

// UpdateOpening
type UpdateOpeningRequest struct {
	Id       int    `json: "id"`
	Role     string `json:"role"`
	Company  string `json: "company"`
	Location string `json: "location"`
	Remote   *bool  `json: "remote"`
	Link     string `json: "link"`
	Salary   int64  `json: "salary"`
}

func (r *UpdateOpeningRequest) Validate() error {

	if r.Id < 1 {
		return fmt.Errorf("resource iD not provided. please include the resource iD in your request and try again.")
	}

	// if any field is provide, valaidation is truthy
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil ||
		r.Link != "" || r.Salary > 0 {
		return nil
	}

	return fmt.Errorf("at least one valid field must be provide")
}
