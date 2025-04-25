package customer

import (
	"testing"

	"github.com/labstack/echo/v4"
)

func TestRegister(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("panic: %v", r)
		}
	}()

	Register(
		echo.New().Group("/api"),
		NewCustomerHandlerCreate(nil),
		NewCustomerHandlerList(nil),
	)
}
