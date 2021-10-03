package Handler

import (
	auth "cager/App/auth"
	"cager/App/category"
	"cager/App/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryService category.Service
	authService     auth.Service
}

func NewCategoryHandler(categoryService category.Service, authService auth.Service) *categoryHandler {
	return &categoryHandler{categoryService, authService}
}

func (h *categoryHandler) RegisterCategory(c *gin.Context) {

	var input category.RegisterCategoryInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Register category failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	_, err = h.categoryService.RegisterCategory(input)

	if err != nil {
		response := helper.APIResponse("Register category failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Category has been registered", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) UpdateCategory(c *gin.Context) {
	var input category.UpdateCategoryInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Update category failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}
	_, err = h.categoryService.UpdateCategory(input)
	if err != nil {
		response := helper.APIResponse("Update category failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Update has been updated", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)

}

func (h *categoryHandler) DeleteCategory(c *gin.Context) {
	var input category.DeleteCategoryInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Delete category failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	del, err := h.categoryService.DeleteCategory(input)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse("Please try again", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if del != true {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse("Please try again", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Delete category", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) GetCategoryById(c *gin.Context) {
	var input category.GetCategoryDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	categoryDetail, err := h.categoryService.GetCategoryById(input)

	if err != nil {
		response := helper.APIResponse("Failed to get detail of category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if categoryDetail.ID == 0 {
		response := helper.APIResponse("Failed to get detail of category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := category.FormatCategory(categoryDetail)
	response := helper.APIResponse("Category detail", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) GetCategoryBySlug(c *gin.Context) {
	var input category.GetCategorySlugInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	categoryDetail, err := h.categoryService.GetCategoryBySlug(input)

	if err != nil {
		response := helper.APIResponse("Failed to get detail of category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if categoryDetail.ID == 0 {
		response := helper.APIResponse("Failed to get detail of category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := category.FormatCategory(categoryDetail)
	response := helper.APIResponse("Category detail", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) GetAllCategory(c *gin.Context) {
	categories, err := h.categoryService.FindAllCategory()

	if err != nil {
		response := helper.APIResponse("Failed to get detail of category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Category detail", http.StatusOK, "success", categories)
	c.JSON(http.StatusOK, response)

}
