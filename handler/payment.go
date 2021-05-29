package handler

import (
	"cager/helper"
	"cager/payment"

	"net/http"

	"github.com/gin-gonic/gin"
)

type paymentHandler struct {
	paymentService payment.Service
}

func NewPaymentHandler(paymentService payment.Service) *paymentHandler {
	return &paymentHandler{paymentService}
}

func (h *paymentHandler) Index(c *gin.Context) {
	payments, err := h.paymentService.GetAllPayment()
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Payment failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		// return payments, err
	}
	response := helper.APIResponse("Payment registered", http.StatusOK, "success", payments)
	c.JSON(http.StatusOK, response)

	// return payments, nil
}

func (h *paymentHandler) RegisterPayment(c *gin.Context) {

	var input payment.RegisterPaymentInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Register payment failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.paymentService.RegisterPayment(input)

	if err != nil {
		response := helper.APIResponse("Register payment failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Payment registered", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
