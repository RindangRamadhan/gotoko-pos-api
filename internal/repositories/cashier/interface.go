package cashier

import (
	"context"
	"gotoko-pos-api/common/response"
	"gotoko-pos-api/internal/entities"
)

//go:generate mockgen -destination=mock/cashier.go -package=mock gotoko-pos-api/internal/repositories/marketplace/cashier IRepo
type IRepo interface {
	GetCashier(ctx context.Context, req entities.CashierFilter) ([]entities.Cashier, response.MetaTpl, error)
	GetCashierById(ctx context.Context, id int64) (entities.Cashier, error)
	CreateCashier(ctx context.Context, req entities.Cashier) (entities.Cashier, error)
	UpdateCashier(ctx context.Context, req entities.Cashier) error
	DeleteCashier(ctx context.Context, id int64) error
}
