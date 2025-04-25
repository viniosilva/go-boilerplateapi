package dto_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/viniosilva/go-boilerplateapi/internal/application/customer/dto"
	"github.com/viniosilva/go-boilerplateapi/internal/domain/customer"
)

func TestFromEntity(t *testing.T) {
	t.Run("should map customer entity to dto successfully", func(t *testing.T) {
		entity := &customer.Customer{
			ID:        1,
			FirstName: "Jane",
			LastName:  "Smith",
			Phone:     "00987654321",
		}

		expected := &dto.Customer{
			ID:        entity.ID,
			FirstName: entity.FirstName,
			LastName:  entity.LastName,
			Phone:     entity.Phone,
		}

		got := dto.FromEntity(entity)
		assert.Equal(t, expected, got)
	})
}

func BenchmarkFromEntity(b *testing.B) {
	entity := &customer.Customer{
		ID:        1,
		FirstName: "Jane",
		LastName:  "Smith",
		Phone:     "00987654321",
	}

	for b.Loop() {
		dto.FromEntity(entity)
	}
}
