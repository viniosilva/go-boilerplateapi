package usecase

import (
	"context"

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
}

func NewCustomerUseCaseCreate(repo customer.CustomerRepository) *CustomerUseCaseCreate {
	meter := otel.Meter("internal/application/customer/usecase/create_customer_usecase")
	counter, _ := meter.Int64Counter("customer_created_total")

	return &CustomerUseCaseCreate{
		repo:    repo,
		tracer:  otel.Tracer("internal/application/customer/usecase/create_customer_usecase"),
		counter: counter,
	}
}

func (uc *CustomerUseCaseCreate) Execute(ctx context.Context, input dto.CreateCustomerInput) (*dto.Customer, error) {
	ctx, span := uc.tracer.Start(ctx, "CustomerUseCaseCreate.Execute")
	defer span.End()

	customer := input.ToEntity()
	if err := uc.repo.Save(ctx, customer); err != nil {
		return nil, err
	}
	uc.counter.Add(ctx, 1)

	return dto.FromEntity(customer), nil
}
