package usecase

import (
	"Customer-RestFull_Api/domain/entity"
	"Customer-RestFull_Api/domain/repository"
	"Customer-RestFull_Api/internal/delivery/request"
	"time"
)

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) GetAllCustomer() ([]entity.Customers, error) {
	customers, err := s.repository.GetAllCustomer()
	return customers, err

}

func (s *service) GetCustomerById(ID int) (entity.Customers, error) {
	customer, err := s.repository.GetCustomerById(ID)
	return customer, err
}

func (s *service) CreateCustomer(customerRequest request.CustomerRequest) (entity.Customers, error) {

	//name, err := string(customerRequest.Name)
	//address, err := customerRequest.Address.String()
	//phoneNumber, err := customerRequest.PhoneNumber.Int64()
	//created_at, err := customerRequest.Created_At.Float64()
	//update_at, err := customerRequest.Update_At.Int64()
	//delete_at, err := customerRequest.Delete_At.Int64()

	customer := entity.Customers{
		Name:        customerRequest.Name,
		Address:     customerRequest.Address,
		Phonenumber: 0,
		Email:       customerRequest.Email,
		Created_at:  time.Time{},
		Update_at:   time.Time{},
		Delete_at:   time.Time{},
	}

	NewCustomer, err := s.repository.CreateCustomer(customer)
	return NewCustomer, err
}

func (s *service) UpdateCustomer(Id int, customerRequest request.CustomerRequest) (entity.Customers, error) {
	customer, err := s.repository.GetCustomerById(Id)

	phonenumber, _ := customerRequest.PhoneNumber.Int64()

	customer.Name = customerRequest.Name
	customer.Address = customerRequest.Address
	customer.Phonenumber = int(phonenumber)
	customer.Email = customerRequest.Email

	NewCustomer, err := s.UpdateCustomer(customer)
	return NewCustomer, err

}

func (s *service) Delete(Id int) (entity.Customers, error) {
	customer, err := s.repository.GetCustomerById(Id)
	NewCustomer, err := s.repository.Delete(customer)
	return NewCustomer, err

}
