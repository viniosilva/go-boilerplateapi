package customer

import (
	"context"

	"github.com/viniosilva/go-boilerplateapi/pkg/pagination"
)

type CustomerRepository interface {
	Save(ctx context.Context, customer *Customer) error
	List(ctx context.Context, params pagination.Params) (pagination.Pagination[Customer], error)
}
