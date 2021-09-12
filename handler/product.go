package handler

import (
	"cager/helper"
	"cager/product"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService product.Service
}

func NewProductHandler(productService product.Service) *productHandler {
	return &productHandler{productService}
}

// CREATE PRODUCT BY NAME
func (h *productHandler) CreateProductName(c *gin.Context) {
	var input product.CreateProductByName

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Validation Product Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	find, err := h.productService.FindProductBySlug(input.Slug)

	if err != nil {
		response := helper.APIResponse("Create Product failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if find != true {
		response := helper.APIResponse("Change With Other Title", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	query, err := h.productService.CreateProductByName(input)
	if err != nil {
		response := helper.APIResponse("Create Product failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := product.FormatProduct(query)
	response := helper.APIResponse("Create Products", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// ALL PRODUCT
func (h *productHandler) GetAllProduct(c *gin.Context) {
	products, err := h.productService.FindAllProductService()
	if err != nil {
		response := helper.APIResponse("Find Products failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := product.FormatProducts(products)
	response := helper.APIResponse("Get All Products", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// SAVE PRODUCT
func (h *productHandler) UpdateProduct(c *gin.Context) {
	var input product.UpdateProductInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Validation Product Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	_, err = h.productService.UpdateProductService(input)
	if err != nil {
		response := helper.APIResponse("Update Products failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Product Update", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)

}

func (h *productHandler) CreateSlider(c *gin.Context) {
	// is_primary set in front end
	// validation input
	// create slider

	var input product.CreateSliderInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Validation Slider Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	slider, err := h.productService.CreateSliderService(input)
	if err != nil {
		response := helper.APIResponse("Created Slider Failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := product.FormatSlider(slider)
	response := helper.APIResponse("Slider Created", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *productHandler) DelSlider(c *gin.Context) {
	var input product.DelSliderInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Validation Slider Failed", http.StatusUnprocessableEntity, "errors", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.productService.DelSliderService(input.ID)

	if err != nil {
		response := helper.APIResponse("Del Slider Failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Slider Deleted", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

// ========= DISCOUNT
func (h *productHandler) CreateDiscount(c *gin.Context) {
	var input product.CreateDiscountInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Validation Discount Failed", http.StatusUnprocessableEntity, "errors", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	discount, err := h.productService.CreateDiscountService(input)
	if err != nil {
		response := helper.APIResponse("Create Discount Failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := product.FormatDiscount(discount)
	response := helper.APIResponse("Discount Created", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}
