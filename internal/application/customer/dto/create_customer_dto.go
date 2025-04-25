package dto

import "github.com/viniosilva/go-boilerplateapi/internal/domain/customer"

type CreateCustomerInput struct {
	FirstName string `json:"first_name" example:"John" validate:"required"`
	LastName  string `json:"last_name" example:"Doe" validate:"required"`
	Phone     string `json:"phone" example:"00123456789" validate:"required"`
}

func (i *CreateCustomerInput) ToEntity() *customer.Customer {
	return customer.NewCustomer(i.FirstName, i.LastName, i.Phone)
}
