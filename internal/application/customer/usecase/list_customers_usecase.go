package usecase

import (
	"context"

	"github.com/viniosilva/go-boilerplateapi/internal/application/customer/dto"
	"github.com/viniosilva/go-boilerplateapi/internal/domain/customer"
	"github.com/viniosilva/go-boilerplateapi/pkg/pagination"
)

type CustomersUseCaseList struct {
	repo customer.CustomerRepository
}

func NewCustomersUseCaseList(repo customer.CustomerRepository) *CustomersUseCaseList {
	return &CustomersUseCaseList{
		repo: repo,
	}
}

func (uc *CustomersUseCaseList) Execute(ctx context.Context, params pagination.Params) (pagination.Pagination[dto.Customer], error) {
	result, err := uc.repo.List(ctx, params)
	if err != nil {
		return pagination.Pagination[dto.Customer]{}, err
	}

	data := make([]dto.Customer, len(result.Data))
	for i, c := range result.Data {
		data[i] = *dto.FromEntity(&c)
	}
	res := pagination.CopyMetadata(result, data)

	return res, nil
}
