package getcategorydetail

import (
	"context"
	"gotoko-pos-api/internal/entities"
)

type Inport interface {
	Execute(context.Context, InportRequest) (InportResponse, error)
}

type InportRequest struct {
	CategoryId int64
}

type InportResponse struct {
	CategoryId int64  `json:"categoryId" example:"1"`
	Name       string `json:"name" example:"Kategori 1"`
}

func NewGetCategoryDetailResponse(category entities.Category) InportResponse {
	return InportResponse{
		CategoryId: category.CategoryId,
		Name:       category.Name,
	}
}
