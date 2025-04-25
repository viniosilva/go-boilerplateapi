package customer

import (
	"context"

	domain "github.com/viniosilva/go-boilerplateapi/internal/domain/customer"
	"github.com/viniosilva/go-boilerplateapi/pkg/pagination"

	"gorm.io/gorm"
)

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *customerRepository {
	return &customerRepository{
		db: db,
	}
}

func (r *customerRepository) Save(ctx context.Context, customer *domain.Customer) error {
	return r.db.WithContext(ctx).Save(customer).Error
}

func (r *customerRepository) List(ctx context.Context, params pagination.Params) (pagination.Pagination[domain.Customer], error) {
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
