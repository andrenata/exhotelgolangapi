package handler

import (
	"cager/helper"
	"cager/settings"
	"net/http"

	"github.com/gin-gonic/gin"
)

type settingHandler struct {
	settingService settings.Service
}

func NewSettingHandler(settingService settings.Service) *settingHandler {
	return &settingHandler{settingService}
}

func (h *settingHandler) UpdateSetting(c *gin.Context) {

	var input settings.UpdateSettingInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Update failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	_, err = h.settingService.UpdateSetting(input)

	if err != nil {
		response := helper.APIResponse("Update failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Updated", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *settingHandler) FindByid(c *gin.Context) {

	var input settings.IDSettingInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Find Setting failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	setting, err := h.settingService.FindById(input.ID)

	if err != nil {
		response := helper.APIResponse("Find Setting failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := settings.FormatSetting(setting)
	response := helper.APIResponse("Setting", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
