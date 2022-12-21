package handler

import (
	"Customer-RestFull_Api/domain/entity"
	"Customer-RestFull_Api/domain/usecase"
	"Customer-RestFull_Api/internal/delivery/request"
	"Customer-RestFull_Api/internal/delivery/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	customerService usecase.Service
}

func NewCustomerHandler(customerService usecase.Service) *customerHandler {
	return &customerHandler{customerService}
}

func (h *customerHandler) GetCustomers(c *gin.Context) {
	customers, err := h.customerService.GetAllCustomer()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var customersResponse []response.CustomerResponse

	for _, c := range customers {
		customerResponse := convertToCustomerResponse(c)
		customersResponse = append(customersResponse, customerResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": customersResponse,
	})
}

func (h *customerHandler) GetCustomer(c *gin.Context) {

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	customer1, err := h.customerService.GetCustomerById(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	customerResponse := convertToCustomerResponse(customer1)

	c.JSON(http.StatusOK, gin.H{
		"data": customerResponse,
	})

}

func (h *customerHandler) CreateCustomer(c *gin.Context) {
	var customerRequest request.CustomerRequest

	err := c.ShouldBindJSON(&customerRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		return
	}

	customer, err := h.customerService.CreateCustomer(customerRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": customer,
	})

}

func (h *customerHandler) UpdateCustomer(c *gin.Context) {
	var customerRequest request.CustomerRequest

	err := c.ShouldBindJSON(&customerRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessages := fmt.Sprintf("Error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		return

	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	customer, err := h.customerService.UpdateCustomer(id, customerRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": customer,
	})
}

func (h *customerHandler) DeleteCustomers(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	c, err := h.customerService.Delete(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}
	customerResponse := convertToCustomerResponse(c)
	c.JSON(http.StatusOK, gin.H{
		"data": customerResponse,
	})
}

func convertToCustomerResponse(c entity.Customers) response.CustomerResponse {
	return response.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		PhoneNumber: c.Phonenumber,
		Email:       c.Email,
		Created_At:  c.Created_at,
		Update_At:   c.Update_at,
		Delete_At:   c.Delete_at,
	}
}
