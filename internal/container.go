package internal

import (
	"gotoko-pos-api/internal/pkg"
	"gotoko-pos-api/internal/pkg/env"
	"gotoko-pos-api/internal/usecases/health"
)

type Container struct {
	HealthCheckUsecase health.Inport
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

	return &Container{
		HealthCheckUsecase: health.NewUsecase(dbClient),
	}
}
