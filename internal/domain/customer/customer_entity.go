package customer

type Customer struct {
	ID        int64
	FirstName string
	LastName  string
	Phone     string
}

func NewCustomer(firstName, lastName, phone string) *Customer {
	return &Customer{
		ID:        0,
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
	}
}
