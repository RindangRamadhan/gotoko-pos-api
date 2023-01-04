package getcashierdetail

import (
	"context"
	"gotoko-pos-api/internal/entities"
)

type Inport interface {
	Execute(context.Context, InportRequest) (InportResponse, error)
}

type InportRequest struct {
	CashierId int64
}

type InportResponse struct {
	CashierId int64  `json:"cashierId" example:"1"`
	Name      string `json:"name" example:"Kasir 1"`
}

func NewGetCashierDetailResponse(cashier entities.Cashier) InportResponse {
	return InportResponse{
		CashierId: cashier.CashierId,
		Name:      cashier.Name,
	}
}
