package category

import "github.com/jmoiron/sqlx"

type repo struct {
	mysql *sqlx.DB
}

func NewRepo(mysql *sqlx.DB) IRepo {
	return &repo{
		mysql: mysql,
	}
}
