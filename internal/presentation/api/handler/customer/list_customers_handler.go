package customer

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/viniosilva/go-boilerplateapi/internal/application/customer/usecase"
	"github.com/viniosilva/go-boilerplateapi/pkg/pagination"
)

type CustomerHandlerList struct {
	useCase *usecase.CustomersUseCaseList
}

func NewCustomerHandlerList(useCase *usecase.CustomersUseCaseList) *CustomerHandlerList {
	return &CustomerHandlerList{
		useCase: useCase,
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
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	res, err := h.useCase.Execute(c.Request().Context(), pagination.Params{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
