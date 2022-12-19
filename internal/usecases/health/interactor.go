package health

import (
	"context"
	"gotoko-pos-api/internal/pkg/env"
	"time"

	"github.com/jmoiron/sqlx"
)

type interactor struct {
	mysql *sqlx.DB
}

// NewUsecase --
func NewUsecase(mysql *sqlx.DB) Inport {
	return interactor{
		mysql: mysql,
	}
}

func (i interactor) Execute(ctx context.Context) InportResponse {
	var mysql = true

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if err := i.mysql.PingContext(ctx); err != nil {
		mysql = false
	}

	return InportResponse{
		Name:    "GoToko POS API Service",
		MySQL:   mysql,
		Version: env.Get().ServiceVersion,
	}
}
