package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"strings"
	"time"

	"gotoko-pos-api/common/constant"
	"gotoko-pos-api/common/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
)

const (
	beginTime = "begin"
)

type bodyWriter struct {
	gin.ResponseWriter
	bodyCache *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
	w.bodyCache.Write(b)
	return w.ResponseWriter.Write(b)
}

func TDRLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// skip swagger
		if strings.Contains(c.Request.URL.Path, "docs") {
			return
		}

		reqID := c.Request.Header.Get(constant.HeaderXRequestID)
		if len(reqID) == 0 {
			reqID = uuid.New().String()
		}

		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, constant.ThreadIDKey, reqID)
		ctx = context.WithValue(ctx, beginTime, time.Now())
		c.Request = c.Request.WithContext(ctx)

		c.Writer = &bodyWriter{bodyCache: bytes.NewBufferString(""), ResponseWriter: c.Writer}

		c.Next()

		bw, ok := c.Writer.(*bodyWriter)
		if !ok {
			logger.Fatal(ctx, "the writer was override, can not read bodyCache", nil)
			return
		}

		errMsg, _ := c.Get(constant.ErrorMessageKey)
		threadID := ctx.Value(constant.ThreadIDKey)

		var reqBody interface{}
		reqRaw, _ := c.GetRawData()
		if len(reqRaw) > 0 {
			reqBody = json.RawMessage(reqRaw)
		}

		tdr := logger.LogTDRModel{
			XTime:         time.Now().Format(time.RFC3339),
			Method:        c.Request.Method,
			AppName:       "GoToko-POS-API",
			AppVersion:    "0.0",
			IP:            "localhost",
			Port:          os.Getenv("PORT"),
			SrcIP:         c.ClientIP(),
			Path:          c.Request.URL.String(),
			Header:        getRequestHeaders(c),
			Request:       reqBody,
			Response:      json.RawMessage(bw.bodyCache.Bytes()),
			ResponseCode:  c.Writer.Status(),
			Error:         cast.ToString(errMsg),
			CorrelationID: cast.ToString(threadID),
		}

		beginTime, ok := ctx.Value(beginTime).(time.Time)
		if ok {
			tdr.RespTime = time.Since(beginTime).Milliseconds()
		}

		logger.TDR(tdr)
	}
}

func getRequestHeaders(c *gin.Context) map[string]interface{} {
	headers := map[string]interface{}{}

	for key, val := range c.Request.Header {
		headers[key] = val[0]
	}

	return headers
}
