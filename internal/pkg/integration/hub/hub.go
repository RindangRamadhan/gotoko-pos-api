package hub

import (
	"context"
	"fmt"
	"net/http"

	"gotoko-pos-api/common/httpclient"
	"gotoko-pos-api/common/logger"
)

type (
	Hub struct {
		ID        int     `json:"id"`
		Name      string  `json:"name"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
)

func (g *client) GetHub(ctx context.Context, req GetHubRequest) (GetHubResponse, error) {
	var resp GetHubResponse

	sc, err := httpclient.PerformRequest(ctx, httpclient.Request{
		Client:  g.client,
		Headers: g.header,
		Method:  http.MethodGet,
		Body:    req,
		URL:     g.url + fmt.Sprintf("/hub/%d", req.ID),
		Result:  &resp,
	})

	if err != nil {
		logger.Error(ctx, "got error on request: ", err, logger.GetCallerTrace())
		return resp, err
	}

	return resp, resp.WriteError(ctx, sc)
}
