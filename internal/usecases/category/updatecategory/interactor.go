package updatecategory

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

func (i interactor) Execute(ctx context.Context, req InportRequest) error {
	payload := NewUpdateCategoryRequest(req)

	err := i.categoryRepo.UpdateCategory(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}
