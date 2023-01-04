package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gotoko-pos-api/common/errors"
	"gotoko-pos-api/common/logger"
	"gotoko-pos-api/common/response"
	"gotoko-pos-api/internal/pkg"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

type any interface{}

func BuildPagination(reqSkip, reqLimit int) (int, int, int) {
	skip := reqSkip

	limit := 10
	if reqLimit > 0 {
		limit = reqLimit
	}

	offset := limit * skip

	return skip, limit, offset
}

func BindParam(c *gin.Context, request any) error {
	var err error

	reqType := reflect.TypeOf(request)
	reqVal := reflect.ValueOf(request)

	for i := 0; i < reqType.Elem().NumField(); i++ {
		var (
			field = reqType.Elem().Field(i)

			key   = field.Tag.Get("query")
			value = c.Query(key)
		)

		result := reqVal.Elem().Field(i)

		switch field.Type.Kind() {
		case reflect.Int:
			if value == "" {
				value = "0"
			}

			val, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return errors.NewErr(http.StatusBadRequest, errors.ErrorInvalidParameter, fmt.Sprintf("%s should be integer, got error: %v", key, err))
			}

			if result.CanSet() {
				result.SetInt(val)
			}

		case reflect.String:
			if result.CanSet() {
				result.SetString(value)
			}
		}
	}

	return err
}

func UnmarshalJSON(c *gin.Context, dest interface{}) error {
	if c.Request.Body == nil {
		return nil
	}

	ctx := c.Request.Context()
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error(ctx, "failed read body bytes", err, logger.GetCallerTrace())
		response.WriteError(c, "failed read body bytes", pkg.FormatErrorToCommonError(pkg.ErrJSONParse))
		return err
	}

	if err := json.Unmarshal(bodyBytes, dest); err != nil {
		logger.Error(ctx, "failed unmarshal json", err, logger.GetCallerTrace())
		response.WriteError(c, "failed unmarshal json", pkg.FormatErrorToCommonError(pkg.ErrJSONParse))
		return err
	}

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	return nil
}
