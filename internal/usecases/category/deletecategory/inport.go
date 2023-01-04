package deletecategory

import (
	"context"
)

type Inport interface {
	Execute(context.Context, InportRequest) error
}

type InportRequest struct {
	CategoryId int64 `json:"category_id" validate:"required"`
}
