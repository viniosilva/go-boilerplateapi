package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/viniosilva/go-boilerplateapi/docs"
	"github.com/viniosilva/go-boilerplateapi/internal/container"
	"github.com/viniosilva/go-boilerplateapi/internal/presentation/api/handler/customer"
	"github.com/viniosilva/go-boilerplateapi/internal/presentation/api/middleware"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewServer(di *container.Container, host, port string, timeout time.Duration) *http.Server {
	addr := fmt.Sprintf("%s:%s", host, port)
	docs.SwaggerInfo.Host = addr

	e := echo.New()
	e.HTTPErrorHandler = middleware.ErrorHandler
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	api := e.Group("/api")
	customer.Register(api, di.CustomerHandlerCreate, di.CustomerHandlerList)

	return &http.Server{
		Addr:              addr,
		Handler:           e,
		ReadHeaderTimeout: timeout,
	}
}
