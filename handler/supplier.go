package handler

import (
	"Amz/helper"
	"Amz/supplier"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type supplierHandler struct {
	supplierService supplier.Service
}

func NewSupllierHandler(supplierService supplier.Service) *supplierHandler {
	return &supplierHandler{supplierService}
}

func (h *supplierHandler) RegisterSupplier(c *gin.Context) {

	var input supplier.RegisterSupllierInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Register account supplier failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newSupplier, err := h.supplierService.RegisterSupplier(input)

	if err != nil {
		response := helper.ApiResponse("Register account supplier failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formater := supplier.FormatSupplier(newSupplier)
	response := helper.ApiResponse("Account supplier has been registerd", http.StatusOK, "success", formater)

	c.JSON(http.StatusOK, response)
}

func (h *supplierHandler) GetSuppliers(c *gin.Context) {
	supplierId, _ := strconv.Atoi(c.Query("id"))

	suppplier, err := h.supplierService.GetSuppliers(supplierId)

	if err != nil {
		response := helper.ApiResponse("Error to get suppliers", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("List of suppliers", http.StatusOK, "success", supplier.FormatSuppliers(suppplier))
	c.JSON(http.StatusOK, response)
}
