package customer

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	appDto "github.com/viniosilva/go-boilerplateapi/internal/application/customer/dto"
	"github.com/viniosilva/go-boilerplateapi/internal/application/customer/usecase"
	"github.com/viniosilva/go-boilerplateapi/pkg/httphelper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type CustomerHandlerCreate struct {
	useCase *usecase.CustomerUseCaseCreate
	tracer  trace.Tracer
	name    string
}

func NewCustomerHandlerCreate(useCase *usecase.CustomerUseCaseCreate) *CustomerHandlerCreate {
	return &CustomerHandlerCreate{
		useCase: useCase,
		tracer:  otel.Tracer("Handler"),
		name:    "CustomerHandlerCreate",
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
	ctx, span := h.tracer.Start(c.Request().Context(), fmt.Sprintf("%s.Handle", h.name))
	defer span.End()

	var input appDto.CreateCustomerInput
	if err := httphelper.BindAndValidate(c, &input); err != nil {
		return err
	}

	res, err := h.useCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, res)
}
