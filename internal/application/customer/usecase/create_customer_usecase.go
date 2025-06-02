package usecase

import (
	"context"
	"fmt"

	"github.com/viniosilva/go-boilerplateapi/internal/application/customer/dto"
	"github.com/viniosilva/go-boilerplateapi/internal/domain/customer"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type CustomerUseCaseCreate struct {
	repo    customer.CustomerRepository
	tracer  trace.Tracer
	counter metric.Int64Counter
	name    string
}

func NewCustomerUseCaseCreate(repo customer.CustomerRepository) *CustomerUseCaseCreate {
	meter := otel.Meter("UseCase")
	counter, _ := meter.Int64Counter("customer_created_total")

	return &CustomerUseCaseCreate{
		repo:    repo,
		tracer:  otel.Tracer("UseCase"),
		counter: counter,
		name:    "CustomerUseCaseCreate",
	}
}

func (uc *CustomerUseCaseCreate) Execute(ctx context.Context, input dto.CreateCustomerInput) (*dto.Customer, error) {
	ctx, span := uc.tracer.Start(ctx, fmt.Sprintf("%s.Execute", uc.name))
	defer span.End()

	customer := input.ToEntity()
	if err := uc.repo.Save(ctx, customer); err != nil {
		return nil, err
	}
	uc.counter.Add(ctx, 1)

	return dto.FromEntity(customer), nil
}
