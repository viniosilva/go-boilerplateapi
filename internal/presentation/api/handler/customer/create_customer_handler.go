package customer

import (
	"net/http"

	"github.com/labstack/echo/v4"
	appDto "github.com/viniosilva/go-boilerplateapi/internal/application/customer/dto"
	"github.com/viniosilva/go-boilerplateapi/internal/application/customer/usecase"
	"github.com/viniosilva/go-boilerplateapi/pkg/httphelper"
)

type CustomerHandlerCreate struct {
	useCase *usecase.CustomersUseCaseCreate
}

func NewCustomerHandlerCreate(useCase *usecase.CustomersUseCaseCreate) *CustomerHandlerCreate {
	return &CustomerHandlerCreate{
		useCase: useCase,
	}
}

// CreateCustomer godoc
// @Summary Create a customer
// @Description Create a new customer
// @Tags customers
// @Accept json
// @Produce json
// @Param input body CreateCustomer true "Customer input"
// @Success 201 {object} Customer
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /customers [post]
func (h *CustomerHandlerCreate) Handle(c echo.Context) error {
	var input appDto.CreateCustomerInput
	if err := httphelper.BindAndValidate(c, &input); err != nil {
		return err
	}

	res, err := h.useCase.Execute(c.Request().Context(), input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, res)
}
