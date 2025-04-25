package usecase

import (
	"context"

	"github.com/viniosilva/go-boilerplateapi/internal/application/customer/dto"
	"github.com/viniosilva/go-boilerplateapi/internal/domain/customer"
)

type CustomersUseCaseCreate struct {
	repo customer.CustomerRepository
}

func NewCustomersUseCaseCreate(repo customer.CustomerRepository) *CustomersUseCaseCreate {
	return &CustomersUseCaseCreate{
		repo: repo,
	}
}

func (uc *CustomersUseCaseCreate) Execute(ctx context.Context, input dto.CreateCustomerInput) (*dto.Customer, error) {
	customer := input.ToEntity()
	if err := uc.repo.Save(ctx, customer); err != nil {
		return nil, err
	}

	return dto.FromEntity(customer), nil
}
