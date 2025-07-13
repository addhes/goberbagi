package helper

import (
	"fmt"

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

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			var message string

			switch e.Tag() {
			case "required":
				message = fmt.Sprintf("%s wajib diisi", e.Field())
			case "email":
				message = fmt.Sprintf("%s harus berupa email yang valid", e.Field())
			case "min":
				message = fmt.Sprintf("%s minimal %s karakter", e.Field(), e.Param())
			case "max":
				message = fmt.Sprintf("%s maksimal %s karakter", e.Field(), e.Param())
			default:
				message = fmt.Sprintf("%s tidak valid", e.Field())
			}

			errors = append(errors, message)
		}
	}

	return errors
}