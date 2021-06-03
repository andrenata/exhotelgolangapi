package handler

import (
	"cager/balance"
	"cager/helper"
	"cager/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type balanceHandler struct {
	balanceService balance.Service
}

func NewBalanceHandler(balanceService balance.Service) *balanceHandler {
	return &balanceHandler{balanceService}
}

func (h *balanceHandler) CreateBalance(c *gin.Context) {
	var input balance.InputTopUp

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to create Top Up", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// get from jwt
	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID
	// userId := 1

	newBalance, err := h.balanceService.TopUpBalance(userId, input)

	if err != nil {
		response := helper.APIResponse("Failed to create Top Up", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := balance.FormatBalance(newBalance)
	response := helper.APIResponse("Success to create Top Up", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *balanceHandler) BalanceApprove(c *gin.Context) {
	var input balance.InputTopUpApprove
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Approve Top Up Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	NewBalanceApprove, err := h.balanceService.TopUpApprove(input)
	if err != nil {
		response := helper.APIResponse("Approve Top Up Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Approve Top Up Success", http.StatusOK, "success", NewBalanceApprove)
	c.JSON(http.StatusOK, response)
}
