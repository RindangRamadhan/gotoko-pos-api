package response

import (
	"fmt"
	"net/http"
	"strings"

	"gotoko-pos-api/common/constant"
	"gotoko-pos-api/common/errors"
	"gotoko-pos-api/common/logger"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	commvalidator "gotoko-pos-api/common/validator"
)

type (
	ValidationError struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}

	ValidationErrors []ValidationError

	MetaTpl struct {
		Total int `json:"total"`
		Limit int `json:"limit"`
		Skip  int `json:"skip"`
	}

	Body struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
		Error   interface{} `json:"errors,omitempty"`
	}

	// Used for swagger response body
	BodySuccess struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    interface{} `json:"-"`
		Meta    *MetaTpl    `json:"-"`
	}

	// Used for swagger response body
	BodyFailure struct {
		Success bool              `json:"success"`
		Message string            `json:"message"`
		Errors  []ValidationError `json:"-"`
	}
)

func WriteSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Body{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func WriteError(c *gin.Context, logMessage string, err error) {
	var errVal ValidationErrors

	logger.Error(c, logMessage, err, logger.GetCallerTrace())

	c.Set(constant.ErrorMessageKey, err.Error())

	httpStatusCode := http.StatusInternalServerError
	payload := Body{
		Error:   errors.ErrorCodeGeneralError,
		Success: false,
		Message: fmt.Sprintf("fatal error: %s", err.Error()),
	}

	switch e := err.(type) {
	case validator.ValidationErrors:
		trans, _ := commvalidator.GetTranslator("en")

		for _, v := range e {
			errVal = append(errVal, ValidationError{
				Field:   strings.ToLower(v.Field()),
				Message: strings.Replace(v.Translate(trans), "_", " ", 1),
			})
		}

		payload.Error = errVal
		payload.Message = http.StatusText(http.StatusBadRequest)
	case (*errors.Err):
		payload.Message = e.Error()
		payload.Error = e.GetErrorCode()
		httpStatusCode = e.GetHttpStatusCode()
	}

	c.JSON(httpStatusCode, payload)
}
