package dto

import "github.com/viniosilva/go-boilerplateapi/internal/domain/customer"

type Customer struct {
	ID        int64  `json:"id" example:"1"`
	FirstName string `json:"first_name" example:"John"`
	LastName  string `json:"last_name" example:"Doe"`
	Phone     string `json:"phone" example:"00123456789"`
}

func FromEntity(c *customer.Customer) *Customer {
	return &Customer{
		ID:        c.ID,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Phone:     c.Phone,
	}
}
