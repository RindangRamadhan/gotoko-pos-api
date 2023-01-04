package createcashier

import (
	"context"
	"gotoko-pos-api/internal/repositories/cashier"
)

type interactor struct {
	cashierRepo cashier.IRepo
}

func NewUsecase(cashierRepo cashier.IRepo) Inport {
	return interactor{
		cashierRepo: cashierRepo,
	}
}

func (i interactor) Execute(ctx context.Context, req InportRequest) (InportResponse, error) {
	payload := NewCreateCashierRequest(req)

	cashier, err := i.cashierRepo.CreateCashier(ctx, payload)
	if err != nil {
		return InportResponse{}, err
	}

	return InportResponse{
		Cashier: cashier,
	}, nil
}
