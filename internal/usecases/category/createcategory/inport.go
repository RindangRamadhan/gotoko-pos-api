package createcategory

import (
	"context"
	"gotoko-pos-api/internal/entities"
)

type Inport interface {
	Execute(context.Context, InportRequest) (InportResponse, error)
}

type InportRequest struct {
	Name string `json:"name" validate:"required"`
}

type InportResponse struct {
	entities.Category
}

func NewCreateCategoryRequest(req InportRequest) entities.Category {
	return entities.Category{
		Name: req.Name,
	}
}
