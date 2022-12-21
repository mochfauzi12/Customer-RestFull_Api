package entity

import (
	"errors"
	"time"
)

type Customers struct {
	Id          int
	Name        string
	Address     string
	Phonenumber int
	Email       string
	Created_at  time.Time
	Update_at   time.Time
	Delete_at   time.Time
}

type DTOCustomers struct {
	Id          int
	Name        string
	Address     string
	PhoneNumber int
	Email       string
	Created_At  string
	Update_At   string
	Delete_At   string
}

func NewCustomers(dto DTOCustomers) (*Customers, error) {
	if dto.Name == "" {
		return nil, errors.New("Name not be empty")
	}
	if dto.Address == "" {
		return nil, errors.New("Address not be empty")
	}
	if dto.PhoneNumber == 0 {
		return nil, errors.New("Phone Number not be empty")
	}
	if dto.Email == "" {
		return nil, errors.New("Email not be empty")
	}

	strCreatedAt, _ := time.Parse("2022-01-02", dto.Created_At)
	strUpdateAt, _ := time.Parse("2022-02-03", dto.Update_At)
	strDeleteAt, _ := time.Parse("2022-03-04", dto.Delete_At)

	costumer := &Customers{
		Id:          dto.Id,
		Name:        dto.Name,
		Address:     dto.Address,
		Phonenumber: dto.PhoneNumber,
		Email:       dto.Email,
		Created_at:  strCreatedAt,
		Update_at:   strUpdateAt,
		Delete_at:   strDeleteAt,
	}
	return costumer, nil
}

func (c *Customers) GetIdCustomer() int {
	return c.Id
}

func (c *Customers) GetName() string {
	return c.Name
}
func (c *Customers) GetAddress() string {
	return c.Address
}
func (c *Customers) GetPhoneNumber() int {
	return c.Phonenumber
}
func (c *Customers) GetEmail() string {
	return c.Email
}
func (c *Customers) GetCreatedAt() string {
	return c.Created_at.Format("2022-01-02")
}

func (c *Customers) GetUpdateAt() string {
	return c.Update_at.Format("2022-02-03")
}
func (c *Customers) GetDeleteAt() string {
	return c.Delete_at.Format("2022-03-04")
}
