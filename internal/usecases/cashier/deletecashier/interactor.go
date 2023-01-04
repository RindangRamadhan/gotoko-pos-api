package deletecashier

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
	err := i.cashierRepo.DeleteCashier(ctx, req.CashierId)
	if err != nil {
		return err
	}

	return nil
}
