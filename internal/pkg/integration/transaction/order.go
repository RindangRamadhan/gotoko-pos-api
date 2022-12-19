package transaction

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"gotoko-pos-api/common/httpclient"
	"gotoko-pos-api/common/logger"

	"github.com/guregu/null"
)

type (
	Order struct {
		ID              int       `json:"id"`
		OrderNumber     string    `json:"order_number"`
		CustomerName    string    `json:"customer_name"`
		CustomerAddress string    `json:"customer_address"`
		Status          string    `json:"status"`
		Note            string    `json:"note"`
		OrderTime       time.Time `json:"order_time"`
	}

	OrderDetail struct {
		OrderNumber   string     `json:"order_number"`
		PaymentMethod string     `json:"payment_method"`
		TotalOrder    float64    `json:"total_order"`
		TotalCoin     null.Float `json:"total_coin"`
		ShippingCosts float64    `json:"shipping_costs"`
		Subtotal      float64    `json:"subtotal"`
		Products      []Product  `json:"products"`
		Customer      Customer   `json:"customer"`
	}
)

func (g *client) GetOrders(ctx context.Context, req GetOrderRequest) (GetOrderResponse, error) {
	var resp GetOrderResponse

	sc, err := httpclient.PerformRequest(ctx, httpclient.Request{
		Client:  g.client,
		Headers: g.header,
		Method:  http.MethodPost,
		Body:    req,
		URL:     g.url + "/courier/orders",
		Result:  &resp,
	})

	if err != nil {
		logger.Error(ctx, "got error on request: ", err, logger.GetCallerTrace())
		return resp, err
	}

	return resp, resp.WriteError(ctx, sc)
}

func (g *client) GetOrderDetail(ctx context.Context, req GetOrderDetailRequest) (GetOrderDetailResponse, error) {
	var resp GetOrderDetailResponse

	sc, err := httpclient.PerformRequest(ctx, httpclient.Request{
		Client:  g.client,
		Headers: g.header,
		Method:  http.MethodGet,
		Body:    req,
		URL:     g.url + fmt.Sprintf("/courier/orders/%d/detail", req.ID),
		Result:  &resp,
	})

	if err != nil {
		logger.Error(ctx, "got error on request: ", err, logger.GetCallerTrace())
		return resp, err
	}

	return resp, resp.WriteError(ctx, sc)
}
