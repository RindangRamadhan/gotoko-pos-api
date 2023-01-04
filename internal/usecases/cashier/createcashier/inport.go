package createcashier

import (
	"context"
	"gotoko-pos-api/internal/entities"
)

type Inport interface {
	Execute(context.Context, InportRequest) (InportResponse, error)
}

type InportRequest struct {
	Name     string `json:"name" validate:"required"`
	Passcode string `json:"passcode" validate:"required,len=6"`
}

type InportResponse struct {
	entities.Cashier
}

func NewCreateCashierRequest(req InportRequest) entities.Cashier {
	return entities.Cashier{
		Name:     req.Name,
		Passcode: req.Passcode,
	}
}
