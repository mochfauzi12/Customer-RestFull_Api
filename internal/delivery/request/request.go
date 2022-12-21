package request

import "encoding/json"

type CustomerRequest struct {
	Name        string      `json:"name" binding:"required"`
	Address     string      `json:"address" binding:"required"`
	PhoneNumber json.Number `json:"phone_number" binding:"required"`
	Email       string      `json:"email" binding:"required"`
	Created_At  json.Number `json:"created_at" binding:"required"`
	Update_At   json.Number `json:"update_at" binding:"required"`
	Delete_At   json.Number `json:"delete_at" binding:"required"`
}
