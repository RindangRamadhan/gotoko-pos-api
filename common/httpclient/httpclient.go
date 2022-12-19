package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"gotoko-pos-api/common/constant"
	"gotoko-pos-api/common/logger"

	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/hystrix"
	"github.com/spf13/cast"
)

type Configuration struct {
	Timeout              time.Duration
	CommandName          string
	HystrixTimeout       time.Duration
	MaxConcurrentRequest int
	ErrorPercentTreshold int
	RetryCount           int
}

type Client struct {
	hystrixClient *hystrix.Client
}

func NewClient(config *Configuration) heimdall.Doer {
	return &Client{
		hystrixClient: config.setupHystrix(),
	}
}

func (config *Configuration) setupHystrix() *hystrix.Client {
	httpClient := &http.Client{
		Timeout: config.Timeout,
	}

	return hystrix.NewClient(
		hystrix.WithHTTPTimeout(config.Timeout),
		hystrix.WithCommandName(config.CommandName),
		hystrix.WithHystrixTimeout(config.HystrixTimeout),
		hystrix.WithMaxConcurrentRequests(config.MaxConcurrentRequest),
		hystrix.WithErrorPercentThreshold(config.ErrorPercentTreshold),
		hystrix.WithHTTPClient(httpClient),
		hystrix.WithRetryCount(config.RetryCount),
	)
}

func (client *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := client.hystrixClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type Request struct {
	Client      heimdall.Doer
	Headers     *http.Header
	Method      string
	Body        interface{}
	URL         string
	Result      interface{}
	QueryParams map[string]string
}

func PerformRequest(ctx context.Context, data Request) (int, error) {
	var (
		body []byte
		err  error
	)

	queryParams := url.Values{}
	url, _ := url.Parse(data.URL)

	if len(data.QueryParams) > 0 && data.QueryParams != nil {
		for key, value := range data.QueryParams {
			queryParams.Add(key, value)
		}
		url.RawQuery = queryParams.Encode()
	}

	if data.Body != nil {
		body, err = json.Marshal(data.Body)
		if err != nil {
			return 0, err
		}
	}

	req, err := http.NewRequest(data.Method, url.String(), bytes.NewBuffer(body))
	if err != nil {
		return 0, err
	}

	req = req.WithContext(ctx)
	req.Header = http.Header{}
	if data.Headers != nil {
		req.Header = *data.Headers
	}
	if req.Header.Get("content-type") == "" {
		req.Header.Add("content-type", "application/json")
	}

	// add request id
	if req.Header.Get(constant.HeaderXRequestID) == "" {
		req.Header.Set(constant.HeaderXRequestID, cast.ToString(ctx.Value(constant.ThreadIDKey)))
	}

	tmBeforeRequest := time.Now()

	resp, err := data.Client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("unable connect to the service, got: %s", err.Error())
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, err
	}
	defer resp.Body.Close()

	logger.Info(req.Context(), "http call info",
		logger.Field{Key: "service", Val: "3rd party service"},
		logger.Field{Key: "url", Val: req.URL},
		logger.Field{Key: "headers", Val: req.Header},
		logger.Field{Key: "method", Val: req.Method},
		logger.Field{Key: "rt", Val: time.Since(tmBeforeRequest).Milliseconds()},
		logger.Field{Key: "httpStatus", Val: resp.StatusCode},
		logger.Field{Key: "request", Val: json.RawMessage(body)},
		logger.Field{Key: "response", Val: json.RawMessage(result)},
	)

	if data.Result == nil {
		return resp.StatusCode, nil
	}

	err = json.Unmarshal(result, data.Result)
	if err != nil {
		logger.Error(ctx, "failed to parse json response", err, logger.GetCallerTrace())
		return resp.StatusCode, err
	}

	return resp.StatusCode, nil
}
