package container

import (
	customerUseCase "github.com/viniosilva/go-boilerplateapi/internal/application/customer/usecase"
	customerRepo "github.com/viniosilva/go-boilerplateapi/internal/infrastructure/db/customer"
	customerHandler "github.com/viniosilva/go-boilerplateapi/internal/presentation/api/handler/customer"

	"gorm.io/gorm"
)

type Container struct {
	CustomerHandlerCreate *customerHandler.CustomerHandlerCreate
	CustomerHandlerList   *customerHandler.CustomerHandlerList
}

func New(db *gorm.DB) *Container {
	customerRepo := customerRepo.NewCustomerRepository(db)

	customersUseCaseCreate := customerUseCase.NewCustomersUseCaseCreate(customerRepo)
	customerHandlerCreate := customerHandler.NewCustomerHandlerCreate(customersUseCaseCreate)

	customersUseCaseList := customerUseCase.NewCustomersUseCaseList(customerRepo)
	customerHandlerList := customerHandler.NewCustomerHandlerList(customersUseCaseList)

	return &Container{
		CustomerHandlerCreate: customerHandlerCreate,
		CustomerHandlerList:   customerHandlerList,
	}
}
