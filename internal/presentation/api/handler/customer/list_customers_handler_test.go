package customer_test

import (
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
	"github.com/viniosilva/go-boilerplateapi/pkg/pagination"
	"go.uber.org/mock/gomock"
)

func TestCustomerHandlerList_Handle(t *testing.T) {
	e := echo.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	customerRepositoryMock := mock.NewMockCustomerRepository(ctrl)
	customersUseCaseList := usecase.NewCustomersUseCaseList(customerRepositoryMock)
	h := customer.NewCustomerHandlerList(customersUseCaseList)

	t.Run("should list customers successfully", func(t *testing.T) {
		expectedCode := http.StatusOK
		expectedBody := pagination.Pagination[dto.Customer]{
			Data: []dto.Customer{
				{
					ID:        1,
					FirstName: "John",
					LastName:  "Doe",
					Phone:     "00123456789",
				},
				{
					ID:        2,
					FirstName: "Jane",
					LastName:  "Smith",
					Phone:     "00987654321",
				},
			},
			Total:      2,
			Page:       1,
			Limit:      10,
			TotalPages: 1,
		}

		customerRepositoryMock.EXPECT().
			List(gomock.Any(), gomock.Any()).
			Return(pagination.Pagination[domain.Customer]{
				Data: []domain.Customer{
					{
						ID:        1,
						FirstName: "John",
						LastName:  "Doe",
						Phone:     "00123456789",
					},
					{
						ID:        2,
						FirstName: "Jane",
						LastName:  "Smith",
						Phone:     "00987654321",
					},
				},
				Total:      2,
				Page:       1,
				Limit:      10,
				TotalPages: 1,
			}, nil)

		req := httptest.NewRequest(http.MethodGet, "/api/customers", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := h.Handle(c)
		require.NoError(t, err)

		var got pagination.Pagination[dto.Customer]
		err = json.Unmarshal(rec.Body.Bytes(), &got)
		require.NoError(t, err)

		assert.Equal(t, expectedBody, got)
		assert.Equal(t, expectedCode, rec.Code)
	})

	t.Run("should list empty customers", func(t *testing.T) {
		expectedCode := http.StatusOK
		expectedBody := pagination.Pagination[dto.Customer]{
			Data:       make([]dto.Customer, 0),
			Total:      0,
			Page:       1,
			Limit:      10,
			TotalPages: 0,
		}

		customerRepositoryMock.EXPECT().
			List(gomock.Any(), gomock.Any()).
			Return(pagination.Pagination[domain.Customer]{
				Data:       make([]domain.Customer, 0),
				Total:      0,
				Page:       1,
				Limit:      10,
				TotalPages: 0,
			}, nil)

		req := httptest.NewRequest(http.MethodGet, "/api/customers", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := h.Handle(c)
		require.NoError(t, err)

		var got pagination.Pagination[dto.Customer]
		err = json.Unmarshal(rec.Body.Bytes(), &got)
		require.NoError(t, err)

		assert.Equal(t, expectedBody, got)
		assert.Equal(t, expectedCode, rec.Code)
	})

	t.Run("should throw error on list customers", func(t *testing.T) {
		expectedErr := "error"

		customerRepositoryMock.EXPECT().
			List(gomock.Any(), gomock.Any()).
			Return(pagination.Pagination[domain.Customer]{}, errors.New("error"))

		req := httptest.NewRequest(http.MethodGet, "/api/customers", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := h.Handle(c)

		assert.EqualError(t, err, expectedErr)
	})
}
