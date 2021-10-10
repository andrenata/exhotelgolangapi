package helper

import (
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type WebResponsePagination struct {
	Meta PaginationMeta `json:"meta"`
	Data interface{} `json:"data"`
}

type PaginationMeta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Total  int64 `json:"total"`
}

func APIPagination(message string, code int, status string, total int64, data interface{}) WebResponsePagination {
	meta := PaginationMeta{
		Message: message,
		Code:    code,
		Status:  status,
		Total: total,
	}

	jsonResponsePagination := WebResponsePagination{
		Meta: meta,
		Data: data,
	}

	return jsonResponsePagination
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

type ResponsePagination struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
