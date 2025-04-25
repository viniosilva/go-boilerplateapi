package customer_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/viniosilva/go-boilerplateapi/internal/application/customer/dto"
	"github.com/viniosilva/go-boilerplateapi/internal/application/customer/usecase"
	domain "github.com/viniosilva/go-boilerplateapi/internal/domain/customer"
	"github.com/viniosilva/go-boilerplateapi/internal/presentation/api/handler/customer"
	"github.com/viniosilva/go-boilerplateapi/mock"
	"go.uber.org/mock/gomock"
)

func TestCustomerHandlerCreate_Handle(t *testing.T) {
	e := echo.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	customerRepositoryMock := mock.NewMockCustomerRepository(ctrl)
	customersUseCaseList := usecase.NewCustomersUseCaseCreate(customerRepositoryMock)
	h := customer.NewCustomerHandlerCreate(customersUseCaseList)

	t.Run("should create customer successfully", func(t *testing.T) {
		input := map[string]any{
			"first_name": "John",
			"last_name":  "Doe",
			"phone":      "00123456789",
		}

		expectedCode := http.StatusCreated
		expectedBody := &dto.Customer{
			ID:        1,
			FirstName: input["first_name"].(string),
			LastName:  input["last_name"].(string),
			Phone:     input["phone"].(string),
		}

		body, err := json.Marshal(input)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/customers", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		customerRepositoryMock.EXPECT().
			Save(gomock.Any(), gomock.Any()).
			DoAndReturn(func(ctx context.Context, c *domain.Customer) error {
				c.ID = 1
				return nil
			})

		err = h.Handle(c)
		require.NoError(t, err)

		var got *dto.Customer
		err = json.Unmarshal(rec.Body.Bytes(), &got)
		require.NoError(t, err)

		assert.Equal(t, expectedBody, got)
		assert.Equal(t, expectedCode, rec.Code)
	})

	t.Run("should throw error on bind and validate", func(t *testing.T) {
		input := map[string]any{}

		expectedError := "validation failed"

		body, err := json.Marshal(input)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/customers", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = h.Handle(c)
		assert.EqualError(t, err, expectedError)
	})

	t.Run("should throw error on save customer", func(t *testing.T) {
		input := map[string]any{
			"first_name": "John",
			"last_name":  "Doe",
			"phone":      "00123456789",
		}

		expectedErr := "error"

		body, err := json.Marshal(input)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/customers", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		customerRepositoryMock.EXPECT().
			Save(gomock.Any(), gomock.Any()).
			Return(errors.New("error"))

		err = h.Handle(c)
		assert.EqualError(t, err, expectedErr)
	})
}
