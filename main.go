package main

import (
	"Amz/handler"
	"Amz/refid"
	"Amz/supplier"
	"Amz/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/dbamz?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Koneksi Berhasil")

	userRepository := user.NewRepository(db)
	supplierRepository := supplier.NewRepository(db)
	refidRepository := refid.NewRepository(db)

	userService := user.NewService(userRepository)
	supplierService := supplier.NewService(supplierRepository)
	refidService := refid.NewService(refidRepository)

	userHandler := handler.NewUserHandler(userService)
	supplierHandler := handler.NewSupllierHandler(supplierService)
	refidHandler := handler.NewRefIdHandler(refidService)

	router := gin.Default()

	api := router.Group("api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/supplier", supplierHandler.RegisterSupplier)
	api.POST("/refid", refidHandler.RegisterRefId)
	api.POST("/email_checkers", userHandler.CheckEmail)
	api.GET("/users", userHandler.GetUsers)
	api.GET("/users/:id", userHandler.GetUser)

	router.Run()

}
