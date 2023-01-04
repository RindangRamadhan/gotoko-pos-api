package updatecategory

import (
	"context"
	"gotoko-pos-api/internal/entities"
)

type Inport interface {
	Execute(context.Context, InportRequest) error
}

type InportRequest struct {
	CategoryId int64  `json:"-" validate:"required"`
	Name       string `json:"name" validate:"required"`
}

func NewUpdateCategoryRequest(req InportRequest) entities.Category {
	return entities.Category{
		CategoryId: req.CategoryId,
		Name:       req.Name,
	}
}
