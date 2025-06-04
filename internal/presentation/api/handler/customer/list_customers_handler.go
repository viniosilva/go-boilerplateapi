package customer

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"github.com/viniosilva/go-boilerplateapi/internal/application/customer/usecase"
	"github.com/viniosilva/go-boilerplateapi/pkg/logger"
	"github.com/viniosilva/go-boilerplateapi/pkg/pagination"
)

type CustomerHandlerList struct {
	useCase *usecase.CustomerUseCaseList
	tracer  trace.Tracer
	name    string
}

func NewCustomerHandlerList(useCase *usecase.CustomerUseCaseList) *CustomerHandlerList {
	return &CustomerHandlerList{
		useCase: useCase,
		tracer:  otel.Tracer("internal/presentation/api/handler/customer/list_customers_handler"),
	}
}

// ListCustomers godoc
// @Summary List customers
// @Description Get a paginated list of customers
// @Tags customers
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Success 200 {object} PaginatedCustomer
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /customers [get]
func (h *CustomerHandlerList) Handle(c echo.Context) error {
	ctx, span := h.tracer.Start(c.Request().Context(), "CustomerHandlerList.Handle")
	defer span.End()

	logger.Info(ctx, "CustomerHandlerList.Handle")

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	res, err := h.useCase.Execute(ctx, pagination.Params{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
