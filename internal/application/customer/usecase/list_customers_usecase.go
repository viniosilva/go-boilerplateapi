package usecase

import (
	"context"
	"fmt"

	"github.com/viniosilva/go-boilerplateapi/internal/application/customer/dto"
	"github.com/viniosilva/go-boilerplateapi/internal/domain/customer"
	"github.com/viniosilva/go-boilerplateapi/pkg/pagination"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type CustomerUseCaseList struct {
	repo   customer.CustomerRepository
	tracer trace.Tracer
	name   string
}

func NewCustomerUseCaseList(repo customer.CustomerRepository) *CustomerUseCaseList {
	return &CustomerUseCaseList{
		repo:   repo,
		tracer: otel.Tracer("UseCase"),
		name:   "CustomerUseCaseList",
	}
}

func (uc *CustomerUseCaseList) Execute(ctx context.Context, params pagination.Params) (pagination.Pagination[dto.Customer], error) {
	ctx, span := uc.tracer.Start(ctx, fmt.Sprintf("%s.Execute", uc.name))
	defer span.End()

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
