package handler

import (
	"bytes"
	"cager/helper"
	"cager/pulsa"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type pulsaHandler struct {
	pulsaService pulsa.Service
}

func NewPulsaHandler(pulsaService pulsa.Service) *pulsaHandler {
	return &pulsaHandler{pulsaService}
}

func (h *pulsaHandler) FindByBrand(c *gin.Context) {

	postBody, _ := json.Marshal(map[string]string{
		"cmd":      "prepaid",
		"username": "wenolooeK13W",
		"sign":     "b9baafeded7a9fc27f3f78f79fd8623b",
	})

	responseBody := bytes.NewBuffer(postBody)

	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post("https://api.digiflazz.com/v1/price-list", "application/json", responseBody)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("List Pulsa failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	defer resp.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("List Pulsa failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	sb := string(body)
	response := helper.APIResponse("List pulsa", http.StatusOK, "success", sb)
	c.JSON(http.StatusOK, response)

}
