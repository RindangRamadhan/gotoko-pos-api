package transaction

import "context"

type ITransactionService interface {
	GetOrders(ctx context.Context, req GetOrderRequest) (GetOrderResponse, error)
	GetOrderDetail(ctx context.Context, req GetOrderDetailRequest) (GetOrderDetailResponse, error)
}
