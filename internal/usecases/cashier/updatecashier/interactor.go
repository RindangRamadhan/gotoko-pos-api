package updatecashier

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

func (i interactor) Execute(ctx context.Context, req InportRequest) error {
	payload := NewUpdateCashierRequest(req)

	err := i.cashierRepo.UpdateCashier(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}
