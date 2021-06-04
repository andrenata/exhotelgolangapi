package handler

import (
	"cager/balance"
	"cager/helper"
	"cager/topup"
	"cager/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type balanceHandler struct {
	balanceService balance.Service
	topupService   topup.Service
}

func NewBalanceHandler(balanceService balance.Service, topupService topup.Service) *balanceHandler {
	return &balanceHandler{balanceService, topupService}
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

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	getAmount, err := h.balanceService.FindService(input.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Error while find top up history", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		txHandle.Rollback()
		return
	}

	if getAmount.Status == 1 {
		response := helper.APIResponse("Have top up before", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if input.Secure != "4CCE55_ANDRE_100%" {
		response := helper.APIResponse("Error secure top up history", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err := h.topupService.WithTrx(txHandle).IncrementMoney(uint(getAmount.UserId), float64(getAmount.Amount)); err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Error while incrementing money", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		txHandle.Rollback()
		return
	}

	if err := h.topupService.WithTrx(txHandle).DecrementMoney(uint(input.ID), float64(getAmount.Amount)); err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Error while decrementing money", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		txHandle.Rollback()
		return
	}

	if err := txHandle.Commit().Error; err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("trx commit error:", http.StatusOK, "success", errorMessage)
		c.JSON(http.StatusOK, response)
	}

	response := helper.APIResponse("Approve Top Up Success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
