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
	"go.uber.org/mock/gomock"
)

func TestCustomersUseCaseCreate_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	customerRepositoryMock := mock.NewMockCustomerRepository(ctrl)

	t.Run("should create customer successfully", func(t *testing.T) {
		ctx := context.Background()
		useCase := usecase.NewCustomersUseCaseCreate(customerRepositoryMock)

		input := dto.CreateCustomerInput{
			FirstName: "John",
			LastName:  "Doe",
			Phone:     "1234567890",
		}
		expected := &dto.Customer{
			ID:        1,
			FirstName: input.FirstName,
			LastName:  input.LastName,
			Phone:     input.Phone,
		}

		customerRepositoryMock.EXPECT().
			Save(gomock.Any(), gomock.Any()).
			DoAndReturn(func(ctx context.Context, c *customer.Customer) error {
				c.ID = 1
				return nil
			})

		got, err := useCase.Execute(ctx, input)
		require.NoError(t, err)

		assert.Equal(t, expected, got)
	})

	t.Run("should throw an error on create customer", func(t *testing.T) {
		ctx := context.Background()
		useCase := usecase.NewCustomersUseCaseCreate(customerRepositoryMock)

		input := dto.CreateCustomerInput{
			FirstName: "John",
			LastName:  "Doe",
			Phone:     "1234567890",
		}
		expectedErr := "error creating customer"

		customerRepositoryMock.EXPECT().
			Save(gomock.Any(), gomock.Any()).
			Return(errors.New(expectedErr))

		_, err := useCase.Execute(ctx, input)

		assert.EqualError(t, err, expectedErr)
	})
}

func BenchmarkCustomersUseCaseCreate_Execute(b *testing.B) {
	ctx := context.Background()
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()

	customerRepositoryMock := mock.NewMockCustomerRepository(ctrl)

	useCase := usecase.NewCustomersUseCaseCreate(customerRepositoryMock)

	input := dto.CreateCustomerInput{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "1234567890",
	}

	customerRepositoryMock.EXPECT().
		Save(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, c *customer.Customer) error {
			c.ID = 1
			return nil
		}).AnyTimes()

	for b.Loop() {
		useCase.Execute(ctx, input)
	}
}
