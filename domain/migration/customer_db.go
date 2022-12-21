package migration

import "time"

type Customers struct {
	ID          uint64
	Name        string
	Address     string
	Phonenumber uint64
	Email       string
	Created_At  time.Time
	Update_At   time.Time
	Delete_At   time.Time
}
