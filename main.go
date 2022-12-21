package main

import (
	"Customer-RestFull_Api/domain/migration"
	"Customer-RestFull_Api/internal/handler"
	"Customer-RestFull_Api/internal/repository"
	"Customer-RestFull_Api/internal/usecase"

	//"Customer-RestFull_Api/vendor/github.com/gin-gonic/gin"
	"log"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/web_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database Connection Error Not Found")
	}

	db.AutoMigrate(&migration.Customers{})

	customerRepository := repository.NewRepository(db)

	customerService := usecase.NewService(customerRepository)

	customerHandler := handler.NewCustomerHandler(customerService)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/customers", customerHandler.GetCustomers)

	v1.GET("/customers/:id", customerHandler.GetCustomer)

	v1.POST("/customers", customerHandler.CreateCustomer)

	v1.PUT("/customers/:id", customerHandler.UpdateCustomer)

	v1.DELETE("/customers/:id", customerHandler.DeleteCustomers)

	router.Run(":8080")
}
