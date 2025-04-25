package httphelper

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Code    int               `json:"-"`
	Message string            `json:"message"`
	Details map[string]string `json:"details,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

var validate = validator.New()

func BindAndValidate[T any](c echo.Context, input *T) *ErrorResponse {
	if err := c.Bind(input); err != nil {
		return &ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid request body",
		}
	}

	if err := validate.Struct(input); err != nil {
		if e, ok := err.(validator.ValidationErrors); ok {
			errs := make(map[string]string)
			for _, fieldErr := range e {
				errs[fieldErr.Field()] = fieldErr.Tag()
			}

			return &ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "validation failed",
				Details: errs,
			}
		}

		return &ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "validation error",
		}
	}

	return nil
}
