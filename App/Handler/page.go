package Handler

import (
	"cager/App/helper"
	"cager/App/pages"
	"net/http"

	"github.com/gin-gonic/gin"
)

type pageHandler struct {
	pageService pages.Service
}

func NewPageHandler(pageService pages.Service) *pageHandler {
	return &pageHandler{pageService}
}

func (h *pageHandler) CreatePage(c *gin.Context) {

	var input pages.CreatePageInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Create Page failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	_, err = h.pageService.CreatePage(input)

	if err != nil {
		response := helper.APIResponse("Create Page failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Page has been registered", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *pageHandler) UpdatePage(c *gin.Context) {

	var input pages.UpdatePageInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Update Page failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	_, err = h.pageService.UpdatePage(input)

	if err != nil {
		response := helper.APIResponse("Update Page failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Page has been updated", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *pageHandler) FindByid(c *gin.Context) {

	var input pages.PageIdInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Find Page failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	page, err := h.pageService.FindById(input.ID)

	if err != nil {
		response := helper.APIResponse("Find Page failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := pages.FormatPage(page)
	response := helper.APIResponse("Find Page", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *pageHandler) FindBySlug(c *gin.Context) {

	var input pages.PageSlugInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Find Page failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	page, err := h.pageService.FindBySlug(input.Slug)

	if err != nil {
		response := helper.APIResponse("Find Page failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := pages.FormatPage(page)
	response := helper.APIResponse("Find Page", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *pageHandler) FindAllPage(c *gin.Context) {

	all, err := h.pageService.FindAll()

	if err != nil {
		response := helper.APIResponse("Find Page failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := pages.FormatPages(all)
	response := helper.APIResponse("Find Page", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *pageHandler) DelPage(c *gin.Context) {

	var input pages.PageIdInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Del Page failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	_, err = h.pageService.DelPage(input.ID)

	if err != nil {
		response := helper.APIResponse("Del Page failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Del Page", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
