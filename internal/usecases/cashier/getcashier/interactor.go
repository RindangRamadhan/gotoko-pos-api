package getcashier

import (
	"context"
	"gotoko-pos-api/internal/entities"
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
	payload := entities.CashierFilter{
		Limit: req.Limit,
		Skip:  req.Skip,
	}

	cashiers, meta, err := i.cashierRepo.GetCashier(ctx, payload)
	if err != nil {
		return InportResponse{}, err
	}

	return InportResponse{
		Cashiers: NewGetCashierResponse(cashiers),
		Meta:     meta,
	}, nil
}
