package getcashier

import (
	"context"
	"gotoko-pos-api/common/response"
	"gotoko-pos-api/internal/entities"
)

type Inport interface {
	Execute(context.Context, InportRequest) (InportResponse, error)
}

type InportRequest struct {
	entities.CashierFilter
}

type InportResponse struct {
	Cashiers []GetCashierResponse `json:"cashiers"`
	Meta     response.MetaTpl     `json:"meta"`
}

type (
	GetCashierResponse struct {
		CashierId int64  `json:"cashierId" example:"1"`
		Name      string `json:"name" example:"Kasir 1"`
	}
)

func NewGetCashierResponse(cashiers []entities.Cashier) []GetCashierResponse {
	var resp []GetCashierResponse

	for _, cashier := range cashiers {
		resp = append(resp, GetCashierResponse{
			CashierId: cashier.CashierId,
			Name:      cashier.Name,
		})
	}

	return resp
}
