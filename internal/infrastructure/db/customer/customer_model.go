package customer

import (
	"time"

	domain "github.com/viniosilva/go-boilerplateapi/internal/domain/customer"

	"gorm.io/gorm"
)

type CustomerModel struct {
	ID        int64          `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	FirstName string         `gorm:"column:first_name"`
	LastName  string         `gorm:"column:last_name"`
	Phone     string         `gorm:"column:phone"`
}

func (CustomerModel) TableName() string {
	return "customers"
}

func (m *CustomerModel) ToEntity() *domain.Customer {
	return &domain.Customer{
		ID:        m.ID,
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Phone:     m.Phone,
	}
}

func FromEntity(c *domain.Customer) *CustomerModel {
	return &CustomerModel{
		ID:        c.ID,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Phone:     c.Phone,
	}
}
