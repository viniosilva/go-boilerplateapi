package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/viniosilva/go-boilerplateapi/docs"
	"github.com/viniosilva/go-boilerplateapi/internal/container"
	"github.com/viniosilva/go-boilerplateapi/internal/presentation/api/handler/customer"
	"github.com/viniosilva/go-boilerplateapi/internal/presentation/api/middleware"
	otelecho "go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

func NewServer(di *container.Container, appName, host, port, swaggerAddr string, timeout time.Duration) *http.Server {
	docs.SwaggerInfo.Host = swaggerAddr

	e := echo.New()
	e.Use(otelecho.Middleware(appName))
	e.HTTPErrorHandler = middleware.ErrorHandler
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	api := e.Group("/api")
	customer.Register(api, di.CustomerHandlerCreate, di.CustomerHandlerList)

	addr := fmt.Sprintf("%s:%s", host, port)
	return &http.Server{
		Addr:              addr,
		Handler:           e,
		ReadHeaderTimeout: timeout,
	}
}
