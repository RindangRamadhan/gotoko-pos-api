package updatecashier

import (
	"context"
	"gotoko-pos-api/internal/entities"
)

type Inport interface {
	Execute(context.Context, InportRequest) error
}

type InportRequest struct {
	CashierId int64  `json:"-" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Passcode  string `json:"passcode" validate:"required,len=6"`
}

func NewUpdateCashierRequest(req InportRequest) entities.Cashier {
	return entities.Cashier{
		CashierId: req.CashierId,
		Name:      req.Name,
		Passcode:  req.Passcode,
	}
}
