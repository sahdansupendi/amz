package handler

import (
	"Amz/helper"
	"Amz/refid"
	"github.com/gin-gonic/gin"
	"net/http"
)

type refidHandler struct {
	refidService refid.Service
}

func NewRefIdHandler(refidService refid.Service) *refidHandler {
	return &refidHandler{refidService}
}

func (h *refidHandler) RegisterRefId(c *gin.Context) {

	var input refid.RegisterRefIdInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Register ref Id failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newRefId, err := h.refidService.RegisterRefId(input)

	if err != nil {
		response := helper.ApiResponse("Register ref Id failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formater := refid.FormatRefId(newRefId)
	response := helper.ApiResponse("Ref Id has been registerd", http.StatusOK, "success", formater)

	c.JSON(http.StatusOK, response)
}
