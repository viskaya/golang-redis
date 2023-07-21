package custom_error

import (
	"aegis_test/libs/custom_string"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type customError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c *customError) Error() string {
	return c.Message
}

func BadRequest(message string) *customError {
	return &customError{Code: http.StatusBadRequest, Message: message}
}

func NotFound(message string) *customError {
	return &customError{Code: http.StatusNotFound, Message: message}
}

func InternalError(err error) *customError {
	log.Println(err)
	return &customError{Code: http.StatusInternalServerError, Message: "Oopss!! Internal server error. Please contact Administrator"}
}

func Unauthorized(message string) *customError {
	return &customError{Code: http.StatusUnauthorized, Message: message}
}

func UnavailableService() *customError {
	return &customError{Code: http.StatusServiceUnavailable, Message: "service unavailable"}
}

type CustomParamRequestError struct {
	Code     int               `json:"code"`
	Messages map[string]string `json:"messages"`
}

func (c *CustomParamRequestError) Error() string {
	return "Bad Params"
}

type ErrorParam struct {
	Field   string
	Message string
}

func RequestValidator(request interface{}) *CustomParamRequestError {
	messages := make(map[string]string)

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			message := ""
			if err.Tag() == "gte" {
				message = fmt.Sprintf("must be greater than or equal %s", err.Param())
			} else if err.Tag() == "lte" {
				message = fmt.Sprintf("must be less or equal than %s", err.Param())
			} else if err.Tag() == "email" {
				message = fmt.Sprint("invalid email format. It should be '[userName][@][domainName]'")
			} else if err.Tag() == "e164" {
				message = fmt.Sprint("invalid phone format. It should be '[+][countryCode][number]'")
			} else {
				message = fmt.Sprint("must not be empty")
			}
			field := custom_string.FirstToLower(err.Field())
			messages[field] = message
		}
	}

	return &CustomParamRequestError{Code: http.StatusBadRequest, Messages: messages}
}
