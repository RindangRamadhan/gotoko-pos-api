package category

import (
	"context"
	"gotoko-pos-api/common/response"
	"gotoko-pos-api/internal/entities"
)

//go:generate mockgen -destination=mock/cashier.go -package=mock gotoko-pos-api/internal/repositories/marketplace/cashier IRepo
type IRepo interface {
	GetCategory(ctx context.Context, req entities.CategoryFilter) ([]entities.Category, response.MetaTpl, error)
	GetCategoryById(ctx context.Context, id int64) (entities.Category, error)
	CreateCategory(ctx context.Context, req entities.Category) (entities.Category, error)
	UpdateCategory(ctx context.Context, req entities.Category) error
	DeleteCategory(ctx context.Context, id int64) error
}
