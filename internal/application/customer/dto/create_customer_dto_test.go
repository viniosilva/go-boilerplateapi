package dto_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/viniosilva/go-boilerplateapi/internal/application/customer/dto"
	"github.com/viniosilva/go-boilerplateapi/internal/domain/customer"
)

func TestCreateCustomerInput_ToEntity(t *testing.T) {
	t.Run("should map create customer input to entity successfully", func(t *testing.T) {
		input := dto.CreateCustomerInput{
			FirstName: "John",
			LastName:  "Doe",
			Phone:     "00123456789",
		}

		expected := &customer.Customer{
			ID:        0,
			FirstName: input.FirstName,
			LastName:  input.LastName,
			Phone:     input.Phone,
		}

		got := input.ToEntity()
		assert.Equal(t, expected, got)
	})
}

func BenchmarkCreateCustomerInput_ToEntity(b *testing.B) {
	input := dto.CreateCustomerInput{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "00123456789",
	}

	for b.Loop() {
		input.ToEntity()
	}
}
