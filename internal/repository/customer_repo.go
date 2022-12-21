package repository

import (
	"Customer-RestFull_Api/domain/entity"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllCustomer() ([]entity.Customers, error) {
	var customers []entity.Customers

	err := r.db.Find(&customers).Error

	return customers, err
}

func (r *repository) GetCustomerById(ID int) (entity.Customers, error) {
	var customers entity.Customers

	err := r.db.Find(&customers, ID).Error

	return customers, err
}

func (r *repository) CreateCustomer(customers entity.Customers) (entity.Customers, error) {
	err := r.db.Create(&customers).Error

	return customers, err
}

func (r *repository) UpdateCustomer(customers entity.Customers) (entity.Customers, error) {
	err := r.db.Save(&customers).Error

	return customers, err
}

func (r *repository) DeleteCustomer(customers entity.Customers) (entity.Customers, error) {
	err := r.db.Delete(&customers).Error

	return customers, err
}
