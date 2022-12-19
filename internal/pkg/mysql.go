package pkg

import (
	"fmt"
	"time"

	"gotoko-pos-api/common/database"

	"github.com/jmoiron/sqlx"
)

type MysqlConfig struct {
	Key      string
	Username string
	Password string
	DBName   string
	Host     string
	Port     string

	MaxIdleConn     int
	MaxOpenConn     int
	MaxLifeTimeConn int
}

func NewMysql(config MysqlConfig) *sqlx.DB {
	client := database.NewConfiguration(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	), config.Key)
	err := database.NewMysql(client)
	if err != nil {
		panic(err)
	}

	database.GetSqlxClient(config.Key).SetMaxIdleConns(config.MaxIdleConn)
	database.GetSqlxClient(config.Key).SetMaxOpenConns(config.MaxOpenConn)
	database.GetSqlxClient(config.Key).SetConnMaxLifetime(time.Duration(config.MaxLifeTimeConn) * time.Second)

	return database.GetSqlxClient(config.Key)
}
