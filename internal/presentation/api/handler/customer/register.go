package customer

import (
	"github.com/viniosilva/go-boilerplateapi/internal/application/customer/dto"
	"github.com/viniosilva/go-boilerplateapi/pkg/httphelper"
	"github.com/viniosilva/go-boilerplateapi/pkg/pagination"

	"github.com/labstack/echo/v4"
)

// swag docs
type Customer dto.Customer                             //@name Customer
type CreateCustomer dto.CreateCustomerInput            //@name CreateCustomer
type PaginatedCustomer pagination.Pagination[Customer] //@name PaginatedCustomer
type ErrorResponse httphelper.ErrorResponse            //@name ErrorResponse

func Register(g *echo.Group, create *CustomerHandlerCreate, list *CustomerHandlerList) {
	g.POST("/customers", create.Handle)
	g.GET("/customers", list.Handle)
}
