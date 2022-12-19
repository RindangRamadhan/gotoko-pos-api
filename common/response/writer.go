package response

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	helpervalidator "gotoko-pos-api/common/validator"
)

type (
	ValidationError struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}

	ValidationErrors []ValidationError

	MetaTpl struct {
		Page      int `json:"page"`
		Limit     int `json:"limit"`
		TotalData int `json:"total_data"`
	}

	Body struct {
		Status  bool        `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
		Error   interface{} `json:"errors,omitempty"`
		Meta    *MetaTpl    `json:"meta,omitempty"`
	}

	// Used for swagger response body
	BodySuccess struct {
		Status  bool        `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"-"`
		Meta    *MetaTpl    `json:"-"`
	}

	// Used for swagger response body
	BodyFailure struct {
		Status  bool              `json:"status"`
		Message string            `json:"message"`
		Errors  []ValidationError `json:"-"`
	}
)

func WriteSuccess(c *gin.Context, message string, data interface{}, meta *MetaTpl) {
	c.JSON(http.StatusOK, Body{
		Status:  true,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func WriteError(c *gin.Context, code int, err interface{}) {
	var errors ValidationErrors
	var body = Body{
		Status:  false,
		Message: http.StatusText(code),
	}

	switch e := err.(type) {
	case validator.ValidationErrors:
		trans, _ := helpervalidator.GetTranslator("en")

		for _, v := range e {
			errors = append(errors, ValidationError{
				Field:   strings.ToLower(v.Field()),
				Message: strings.Replace(v.Translate(trans), "_", " ", 1),
			})
		}

		body.Error = errors
	case error:
		body.Message = e.Error()
	default:
		body.Error = err
	}

	c.JSON(code, body)
}
