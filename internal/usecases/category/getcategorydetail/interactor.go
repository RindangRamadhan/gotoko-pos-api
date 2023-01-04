package getcategorydetail

import (
	"context"
	"gotoko-pos-api/internal/repositories/category"
	"log"
)

type interactor struct {
	categoryRepo category.IRepo
}

func NewUsecase(categoryRepo category.IRepo) Inport {
	return interactor{
		categoryRepo: categoryRepo,
	}
}

func (i interactor) Execute(ctx context.Context, req InportRequest) (InportResponse, error) {
	log.Println("req", req)
	category, err := i.categoryRepo.GetCategoryById(ctx, req.CategoryId)
	if err != nil {
		return InportResponse{}, err
	}

	return NewGetCategoryDetailResponse(category), nil
}
