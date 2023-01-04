package deletecategory

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
	err := i.categoryRepo.DeleteCategory(ctx, req.CategoryId)
	if err != nil {
		return err
	}

	return nil
}
