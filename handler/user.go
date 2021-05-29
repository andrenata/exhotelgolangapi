package handler

import (
	"cager/auth"
	"cager/helper"
	"cager/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input
	// map input dari user ke struct RegisterUserInput
	// struct di atas kita passing sebagai parameter service

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// token
	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, token)

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// token
	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, token)
	response := helper.APIResponse("Login success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) ChekEmailAvailability(c *gin.Context) {
	// input email dari user
	// input email di mapping ke struct
	// struct input di passing ke service
	// service akan manggil repository
	// repository akan melakukan query ke database
	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Check email failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.APIResponse("Check email failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"
	metaStatus := "error"

	if isEmailAvailable {
		metaMessage = "Email is available"
		metaStatus = "success"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, metaStatus, data)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) ChekPhoneAvailability(c *gin.Context) {
	var input user.CheckPhoneInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Input check phone failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isPhoneAvailable, err := h.userService.IsPhoneAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.APIResponse("Check phone failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isPhoneAvailable,
	}

	metaMessage := "Phone number has been registered"
	metaStatus := "error"

	if isPhoneAvailable {
		metaMessage = "Phone number is available"
		metaStatus = "success"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, metaStatus, data)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	// input data from user
	// simpan gambarnya di gambar "/images"
	// di service kita panggil repo
	// JWT (sementara hardcode, seakan akan user id = 1)
	// repo ambil data user = 1
	// repo update data user simpan lokasi file
	// c.SaveUploadedFile(file, )
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar", http.StatusUnprocessableEntity, "error", data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// get from jwt
	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	// path := "images/" + file.Filename
	path := fmt.Sprintf("images/%d-%s", userId, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar", http.StatusUnprocessableEntity, "error", data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.userService.SaveAvatar(userId, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Avatar uploaded", http.StatusUnprocessableEntity, "error", data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) HandlerChangePin(c *gin.Context) {
	var input user.ChangePin
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("PIN updated failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	_, err = h.userService.ServiceChangePin(userId, input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("PIN Updated failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{"is_updated": true}
	response := helper.APIResponse("PIN updated", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) HandlerChangePinTemporary(c *gin.Context) {
	var input user.ChangePinTemporary
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Input PIN failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	_, err = h.userService.ServiceChangePinTemporary(userId, input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("PIN Updated failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{"is_updated": true}
	response := helper.APIResponse("PIN updated", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) ChangeName(c *gin.Context) {
	var input user.ChangeNameInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Change name failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	// get from jwt
	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	_, err = h.userService.ServiceChangeName(userId, input)
	if err != nil {
		response := helper.APIResponse("Change name failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Changed name success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) HandlerCheckPin(c *gin.Context) {
	var input user.CheckPin
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("PIN failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	isCheckPin, err := h.userService.ServiceCheckPin(userId, input)
	if err != nil {
		response := helper.APIResponse("Server Error", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_true": isCheckPin,
	}

	metaMessage := "Check is different"
	metaStatus := "error"

	if isCheckPin {
		metaMessage = "PIN success"
		metaStatus = "success"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, metaStatus, data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) HandlerCheckPinTemporary(c *gin.Context) {
	var input user.CheckPin
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("PIN failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	isCheckPin, err := h.userService.ServiceCheckPinTemporary(userId, input)
	if err != nil {
		response := helper.APIResponse("Server Error", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_true": isCheckPin,
	}

	metaMessage := "Check is different"
	metaStatus := "error"

	if isCheckPin {
		metaMessage = "PIN success"
		metaStatus = "success"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, metaStatus, data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) HandlerChangePhoneNumber(c *gin.Context) {
	var input user.InputChangeNumber
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Change phone number failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID
	_, err = h.userService.ServiceChangePhoneNumber(userId, input)
	if err != nil {
		response := helper.APIResponse("Change phone number failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Changed phone number success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) ChangeEmailHandler(c *gin.Context) {
	var input user.ChangeEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Change email failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	_, err = h.userService.ChangeEmailService(userId, input)
	if err != nil {
		response := helper.APIResponse("Change email failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
	}
	response := helper.APIResponse("Change email success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) FetchUser(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(user.User)

	formatter := user.FormatUser(currentUser, "")

	response := helper.APIResponse("Successfuly fetch user data", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
