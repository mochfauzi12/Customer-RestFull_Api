package usecase

import (
	"Customer-RestFull_Api/domain/entity"
	"Customer-RestFull_Api/internal/delivery/request"
)

type Service interface {
	GetAllCustomer() ([]entity.Customers, error)
	GetCustomerById(Id int) (entity.Customers, error)
	CreateCustomer(customerRequest request.CustomerRequest) (entity.Customers, error)
	UpdateCustomer(Id int, customerRequest request.CustomerRequest) (entity.Customers, error)
	Delete(Id int) (entity.Customers, error)
}
