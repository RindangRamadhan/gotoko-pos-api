package getcategory

import (
	"context"
	"gotoko-pos-api/internal/entities"
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
	payload := entities.CategoryFilter{
		Limit: req.Limit,
		Skip:  req.Skip,
	}

	categories, meta, err := i.categoryRepo.GetCategory(ctx, payload)
	if err != nil {
		return InportResponse{}, err
	}

	return InportResponse{
		Categories: NewGetCategoryResponse(categories),
		Meta:       meta,
	}, nil
}
