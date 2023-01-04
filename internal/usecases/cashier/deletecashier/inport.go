package deletecashier

import (
	"context"
)

type Inport interface {
	Execute(context.Context, InportRequest) error
}

type InportRequest struct {
	CashierId int64 `json:"cashier_id" validate:"required"`
}
