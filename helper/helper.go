package helper

import (
	"encoding/json"
	"errors"

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

	jsonRespose := Response{
		Meta: meta,
		Data: data,
	}

	return jsonRespose
}

func FormatValidationError(err error) []string {
	var theErrors []string
	var jsErr *json.UnmarshalTypeError

	if !errors.As(err, &jsErr) {
		for _, e := range err.(validator.ValidationErrors) {
			theErrors = append(theErrors, e.Error())
		}
	} else {
		// add by Me :D
		theErrors = append(theErrors, err.Error())
	}

	return theErrors
}
