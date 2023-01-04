package pkg

import (
	"errors"
	"fmt"
	"net/http"

	commonerror "gotoko-pos-api/common/errors"
)

const (
	ErrCodeDefault = "SERVER-00"

	ErrCodeNotFound   = "DATA-01"
	ErrCodeFatalQuery = "QUERY-01"
	ErrCodePayload    = "PAYLOAD-01"
)

var (
	ErrFatalQuery = errors.New("fatal query error")
	ErrJSONParse  = errors.New("an error occurred in the input parameter")

	ErrCashierNotFound  = errors.New("cashier not found")
	ErrCategoryNotFound = errors.New("category not found")
)

func FormatErrorToCommonError(err error) *commonerror.Err {
	// FATAL QUERY
	ErrInternalFatalQueries := []error{
		ErrFatalQuery,
	}

	if ErrorContains(ErrInternalFatalQueries, err) {
		return commonerror.NewErr(http.StatusInternalServerError, ErrCodeFatalQuery, err.Error())
	}

	// BAD REQUEST
	ErrBadRequest := []error{
		ErrJSONParse,
	}

	if ErrorContains(ErrBadRequest, err) {
		return commonerror.NewErr(http.StatusBadRequest, ErrCodePayload, err.Error())
	}

	// NOT FOUND
	ErrDataNotFound := []error{
		ErrCashierNotFound,
		ErrCategoryNotFound,
	}

	if ErrorContains(ErrDataNotFound, err) {
		return commonerror.NewErr(http.StatusNotFound, ErrCodeNotFound, err.Error())
	}

	// DEFAULT
	return commonerror.NewErr(http.StatusInternalServerError, ErrCodeDefault, fmt.Sprintf("unexpected error from server : %v", err))
}

func ErrorContains(errors []error, err error) bool {
	for _, e := range errors {
		if e == err {
			return true
		}
	}

	return false
}
