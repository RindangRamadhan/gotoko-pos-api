package getcashierdetail

import (
	"context"
	"gotoko-pos-api/internal/repositories/cashier"
	"log"
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
	log.Println("req", req)
	cashier, err := i.cashierRepo.GetCashierById(ctx, req.CashierId)
	if err != nil {
		return InportResponse{}, err
	}

	return NewGetCashierDetailResponse(cashier), nil
}
