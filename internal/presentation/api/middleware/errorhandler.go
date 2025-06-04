package middleware

import (
	"errors"
	"net/http"

	"github.com/viniosilva/go-boilerplateapi/pkg/httphelper"
	"github.com/viniosilva/go-boilerplateapi/pkg/logger"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	errorRes := &httphelper.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "internal server error",
	}

	var echoErr *echo.HTTPError
	if errors.As(err, &echoErr) {
		errorRes.Code = echoErr.Code
		errorRes.Message, _ = echoErr.Message.(string)
	} else if !errors.As(err, &errorRes) {
		logger.Error(c.Request().Context(), err.Error())
	}

	_ = c.JSON(errorRes.Code, errorRes)
}
