package Handler

import (
	"cager/App/category"
	"cager/App/helper"
	product "cager/App/product"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService product.Service
}

func NewProductHandler(productService product.Service) *productHandler {
	return &productHandler{productService}
}

// SLIDER RELATION
func (h *productHandler) CreateSliderRelationHandler(c *gin.Context) {
	var input product.CreateSliderRelationInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Validation Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	check, err := h.productService.CheckSliderRelation(input.ProductID, input.SliderID)
	if err != nil {
		response := helper.APIResponse("Slider failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if !check {
		response := helper.APIResponse("Slider has been selected", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	sliderRelation, err := h.productService.CreateSliderRelationService(input)
	if err != nil {
		response := helper.APIResponse("Create Slider failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := product.FormatSliderRelation(sliderRelation)
	response := helper.APIResponse("Create Products", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *productHandler) GetSliderRelationByProductIDHanlder(c *gin.Context) {
	var input product.IDSliderRelationInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Validation Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	sliderRelation, err := h.productService.GetSliderRelationByIDProductService(input.ID)
	if err != nil {
		response := helper.APIResponse("Get Slider failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := product.FormatSliders(sliderRelation)
	response := helper.APIResponse("Get Slider", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) GetSliderRelationByIDHanlder(c *gin.Context) {
	var input product.IDSliderRelationInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Validation Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	sliderRelation, err := h.productService.GetSliderRelationByIDService(input.ID)
	if err != nil {
		response := helper.APIResponse("Get Slider failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := product.FormatSliderRelation(sliderRelation)
	response := helper.APIResponse("Get Slider", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) DelSliderRelationHanlder(c *gin.Context) {
	var input product.DelSliderProductInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Validation Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	_, err = h.productService.DelSliderRelation(input.SliderID, input.ProductID)
	if err != nil {
		response := helper.APIResponse("Del Slider failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Del Slider", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
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

	if !find {
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

func (h *productHandler) FindProductByIDHandler(c *gin.Context) {
	var input product.FindProductByIdInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Validation Product Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	find, err := h.productService.FindProductById(input.ID)
	if err != nil {
		response := helper.APIResponse("Find Product failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := product.FormatProduct(find)
	response := helper.APIResponse("Get Product", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) GetAllSliderHanlder(c *gin.Context) {
	sliders, err := h.productService.FindAllSliderService()
	if err != nil {
		response := helper.APIResponse("Find Sliders failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := product.FormatSliders(sliders)
	response := helper.APIResponse("Get All Products", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// CATEGORY RELATION
func (h *productHandler) CreateCategoryRelationHandler(c *gin.Context) {
	var input product.CreateCategoryRelationInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Validation Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	check, err := h.productService.CheckCategoryRelation(input.ProductID, input.CategoryID)
	if err != nil {
		response := helper.APIResponse("Create failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if !check {
		response := helper.APIResponse("Category has been selected", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	query, err := h.productService.CreateCategoryRelation(input)
	if err != nil {
		response := helper.APIResponse("Create failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := product.FormatCategoryRelation(query)
	response := helper.APIResponse("Create Success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *productHandler) DelCategoryRelationHandler(c *gin.Context) {
	var input product.CreateCategoryRelationInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Validation Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	_, err = h.productService.DelCategoryRelation(input.ProductID, input.CategoryID)
	if err != nil {
		response := helper.APIResponse("Delete failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Delete Success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) FindCategoryRelationHandler(c *gin.Context) {
	var input product.IDSliderRelationInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Validation Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	find, err := h.productService.FindCategoryRelation(input.ID)
	if err != nil {
		response := helper.APIResponse("Find failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := category.FormatCategories(find)
	response := helper.APIResponse("Find Success", http.StatusOK, "success", formatter)
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

	file, err := c.FormFile("file")

	if err != nil {
		response := helper.APIResponse("Created Slider Failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Set Folder untuk menyimpan filenya
	path := "storage/" + file.Filename
	if err := c.SaveUploadedFile(file, path); err != nil {
		response := helper.APIResponse("Created Slider Failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.productService.CreateSliderService(file.Filename)
	if err != nil {
		response := helper.APIResponse("Created Slider Failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Slider Created", http.StatusOK, "success", nil)
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

func (h *productHandler) DelProduct(c *gin.Context) {
	var input product.DelProductInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Validation Product Failed", http.StatusUnprocessableEntity, "errors", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.productService.DelProductService(input.ID)

	if err != nil {
		response := helper.APIResponse("Del Product Failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Product Deleted", http.StatusOK, "success", nil)
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
