package hub

import (
	"context"
	"net/http"
	"time"

	"gotoko-pos-api/common/httpclient"
	"gotoko-pos-api/common/logger"

	"github.com/gojek/heimdall"
)

type client struct {
	url    string
	client heimdall.Doer
	header *http.Header
}

func NewClient(url string, token string) IHubService {
	header := &http.Header{}
	header.Set("Authorization", token)

	return &client{
		url: url,
		client: httpclient.NewClient(&httpclient.Configuration{
			Timeout:              120 * time.Second,
			CommandName:          "hub_service_http_request",
			HystrixTimeout:       60 * time.Second,
			MaxConcurrentRequest: 30,
			ErrorPercentTreshold: 20,
		}),
		header: header,
	}
}

func (b BaseResponse) WriteError(ctx context.Context, httpCode int) error {
	if httpCode <= http.StatusAccepted {
		return nil
	}

	logger.Error(ctx, "got error from hub service: "+b.Message, nil, logger.GetCallerTrace())
	return ErrrGeneralServiceError
}
