package createcategory

import (
	"context"
	"gotoko-pos-api/internal/repositories/category"
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
	payload := NewCreateCategoryRequest(req)

	category, err := i.categoryRepo.CreateCategory(ctx, payload)
	if err != nil {
		return InportResponse{}, err
	}

	return InportResponse{
		Category: category,
	}, nil
}
