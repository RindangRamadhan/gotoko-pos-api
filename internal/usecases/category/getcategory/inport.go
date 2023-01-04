package getcategory

import (
	"context"
	"gotoko-pos-api/common/response"
	"gotoko-pos-api/internal/entities"
)

type Inport interface {
	Execute(context.Context, InportRequest) (InportResponse, error)
}

type InportRequest struct {
	entities.CategoryFilter
}

type InportResponse struct {
	Categories []GetCategoryResponse `json:"categories"`
	Meta       response.MetaTpl      `json:"meta"`
}

type (
	GetCategoryResponse struct {
		CategoryId int64  `json:"categoryId" example:"1"`
		Name       string `json:"name" example:"Kasir 1"`
	}
)

func NewGetCategoryResponse(categories []entities.Category) []GetCategoryResponse {
	var resp []GetCategoryResponse

	for _, cashier := range categories {
		resp = append(resp, GetCategoryResponse{
			CategoryId: cashier.CategoryId,
			Name:       cashier.Name,
		})
	}

	return resp
}
