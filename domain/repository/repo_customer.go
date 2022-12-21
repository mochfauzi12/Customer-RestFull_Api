package repository

import "Customer-RestFull_Api/domain/entity"

type Repository interface {
	GetAllCustomer() ([]entity.Customers, error)
	GetCustomerById(Id int) (entity.Customers, error)
	CreateCustomer(customer entity.Customers) (entity.Customers, error)
	UpdateCustomer(customer entity.Customers) (entity.Customers, error)
	DeleteCustomer(Customer entity.Customers) (entity.Customers, error)
}
