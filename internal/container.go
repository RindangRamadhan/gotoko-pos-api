package internal

import (
	"gotoko-pos-api/internal/pkg"
	"gotoko-pos-api/internal/pkg/env"
	"gotoko-pos-api/internal/repositories/cashier"
	"gotoko-pos-api/internal/usecases/cashier/createcashier"
	"gotoko-pos-api/internal/usecases/cashier/deletecashier"
	"gotoko-pos-api/internal/usecases/cashier/getcashier"
	"gotoko-pos-api/internal/usecases/cashier/getcashierdetail"
	"gotoko-pos-api/internal/usecases/cashier/updatecashier"
	"gotoko-pos-api/internal/usecases/health"
)

type Container struct {
	HealthCheckUsecase health.Inport

	GetCashierUsecase       getcashier.Inport
	GetCashierDetailUsecase getcashierdetail.Inport
	CreateCashierUsecase    createcashier.Inport
	UpdateCashierUsecase    updatecashier.Inport
	DeleteCashierUsecase    deletecashier.Inport
}

func NewContainer() *Container {
	dbClient := pkg.NewMysql(pkg.MysqlConfig{
		Key:             "gotoko_pos_sqlx",
		Username:        env.Get().DBUsername,
		Password:        env.Get().DBPassword,
		DBName:          env.Get().DBName,
		Host:            env.Get().DBHost,
		Port:            env.Get().DBPort,
		MaxIdleConn:     env.Get().DBMaxIdleConn,
		MaxOpenConn:     env.Get().DBMaxOpenConn,
		MaxLifeTimeConn: env.Get().DBMaxLifeTimeConn,
	})

	// Repo for query related
	cashierRepo := cashier.NewRepo(dbClient)

	return &Container{
		HealthCheckUsecase: health.NewUsecase(dbClient),

		GetCashierUsecase:       getcashier.NewUsecase(cashierRepo),
		GetCashierDetailUsecase: getcashierdetail.NewUsecase(cashierRepo),
		CreateCashierUsecase:    createcashier.NewUsecase(cashierRepo),
		UpdateCashierUsecase:    updatecashier.NewUsecase(cashierRepo),
		DeleteCashierUsecase:    deletecashier.NewUsecase(cashierRepo),
	}
}
