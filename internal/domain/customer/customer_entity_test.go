package customer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCustomer(t *testing.T) {
	t.Run("should create customer sucessfully", func(t *testing.T) {
		expected := &Customer{
			ID:        0,
			FirstName: "John",
			LastName:  "Doe",
			Phone:     "00123456789",
		}
		got := NewCustomer(expected.FirstName, expected.LastName, expected.Phone)

		assert.Equal(t, expected, got)
	})
}

func BenchmarkNewCustomer(b *testing.B) {
	input := &Customer{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "00123456789",
	}

	for b.Loop() {
		NewCustomer(input.FirstName, input.LastName, input.Phone)
	}
}
