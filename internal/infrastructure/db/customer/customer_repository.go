package customer

import (
	"context"

	domain "github.com/viniosilva/go-boilerplateapi/internal/domain/customer"
	"github.com/viniosilva/go-boilerplateapi/pkg/pagination"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	db     *gorm.DB
	tracer trace.Tracer
	name   string
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		db:     db,
		tracer: otel.Tracer("internal/infrastructure/db/customer/customer_repository"),
	}
}

func (r *CustomerRepository) Save(ctx context.Context, customer *domain.Customer) error {
	ctx, span := r.tracer.Start(ctx, "CustomerRepository.Save")
	defer span.End()

	return r.db.WithContext(ctx).Save(customer).Error
}

func (r *CustomerRepository) List(ctx context.Context, params pagination.Params) (pagination.Pagination[domain.Customer], error) {
	ctx, span := r.tracer.Start(ctx, "CustomerRepository.List")
	defer span.End()

	db := r.db.WithContext(ctx)

	params.Normalize()
	res := pagination.Pagination[domain.Customer]{
		Page:  params.Page,
		Limit: params.Limit,
	}

	if err := db.Model(&CustomerModel{}).Count(&res.Total).Error; err != nil {
		return res, err
	}
	res.SetTotalPages()

	var models []CustomerModel
	if err := db.
		Limit(params.Limit).
		Offset(params.CalculateOffset()).
		Find(&models).Error; err != nil {
		return res, err
	}

	res.Data = make([]domain.Customer, len(models))
	for i, m := range models {
		res.Data[i] = *m.ToEntity()
	}

	return res, nil
}
