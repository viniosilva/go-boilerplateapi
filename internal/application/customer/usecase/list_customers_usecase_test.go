package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/viniosilva/go-boilerplateapi/internal/application/customer/dto"
	"github.com/viniosilva/go-boilerplateapi/internal/application/customer/usecase"
	"github.com/viniosilva/go-boilerplateapi/internal/domain/customer"
	"github.com/viniosilva/go-boilerplateapi/mock"
	"github.com/viniosilva/go-boilerplateapi/pkg/pagination"
	"go.uber.org/mock/gomock"
)

func TestCustomersUseCaseList_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	customerRepositoryMock := mock.NewMockCustomerRepository(ctrl)

	t.Run("should list customers successfully", func(t *testing.T) {
		ctx := context.Background()
		useCase := usecase.NewCustomersUseCaseList(customerRepositoryMock)

		expected := pagination.Pagination[dto.Customer]{
			Data: []dto.Customer{
				{
					ID:        1,
					FirstName: "John",
					LastName:  "Doe",
					Phone:     "00123456789",
				},
			},
			Total: 1, Page: 1, Limit: 10, TotalPages: 1,
		}

		customerRepositoryMock.EXPECT().
			List(gomock.Any(), gomock.Any()).
			Return(pagination.Pagination[customer.Customer]{
				Data: []customer.Customer{
					{
						ID:        expected.Data[0].ID,
						FirstName: expected.Data[0].FirstName,
						LastName:  expected.Data[0].LastName,
						Phone:     expected.Data[0].Phone,
					},
				},
				Total: 1, Page: 1, Limit: 10, TotalPages: 1,
			}, nil)

		got, err := useCase.Execute(ctx, pagination.Params{
			Page:  expected.Page,
			Limit: expected.Limit,
		})
		require.NoError(t, err)

		assert.Equal(t, expected, got)
	})

	t.Run("should throw an error on create customer", func(t *testing.T) {
		ctx := context.Background()
		useCase := usecase.NewCustomersUseCaseList(customerRepositoryMock)

		expectedErr := "error creating customer"

		customerRepositoryMock.EXPECT().
			List(gomock.Any(), gomock.Any()).
			Return(pagination.Pagination[customer.Customer]{}, errors.New(expectedErr))

		_, err := useCase.Execute(ctx, pagination.Params{Page: 1, Limit: 10})

		assert.EqualError(t, err, expectedErr)
	})
}

func BenchmarkCustomersUseCaseList_Execute(b *testing.B) {
	ctx := context.Background()
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()

	customerRepositoryMock := mock.NewMockCustomerRepository(ctrl)

	useCase := usecase.NewCustomersUseCaseList(customerRepositoryMock)

	params := pagination.Params{Page: 1, Limit: 10}

	customerRepositoryMock.EXPECT().
		List(gomock.Any(), gomock.Any()).
		Return(pagination.Pagination[customer.Customer]{
			Data: []customer.Customer{
				{
					ID:        1,
					FirstName: "John",
					LastName:  "Doe",
					Phone:     "00123456789",
				},
			},
			Total: 1, Page: 1, Limit: 10, TotalPages: 1,
		}, nil).
		AnyTimes()

	for b.Loop() {
		useCase.Execute(ctx, params)
	}
}
