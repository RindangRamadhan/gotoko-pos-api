package cashier

import (
	"context"
	"database/sql"
	"gotoko-pos-api/common/logger"
	"gotoko-pos-api/common/request"
	"gotoko-pos-api/common/response"
	"gotoko-pos-api/internal/entities"
	"gotoko-pos-api/internal/pkg"
	"time"

	dbtx "gotoko-pos-api/common/database"

	"github.com/jmoiron/sqlx"
)

func (r *repo) GetCashier(ctx context.Context, req entities.CashierFilter) ([]entities.Cashier, response.MetaTpl, error) {
	var (
		pagination response.MetaTpl
		result     struct {
			Rows     []entities.Cashier
			TotalRow int `json:"total_row"`
		}
	)

	// Build pagination
	skip, limit, offset := request.BuildPagination(req.Skip, req.Limit)

	query := `
		SELECT
			cashier_id, name, COUNT(cashier_id) OVER() AS row_count
		FROM cashiers
		ORDER BY cashier_id DESC
		LIMIT ? OFFSET ?
	`
	err := r.mysql.SelectContext(ctx, &result.Rows, query, limit, offset)
	if err != nil {
		logger.Error(ctx, "failed to get cashiers", err, logger.GetCallerTrace())
		return result.Rows, pagination, pkg.FormatErrorToCommonError(pkg.ErrFatalQuery)
	}

	if len(result.Rows) > 0 {
		result.TotalRow = result.Rows[0].RowCount
	}

	pagination.Skip = skip
	pagination.Limit = limit
	pagination.Total = result.TotalRow

	return result.Rows, pagination, nil
}

func (r *repo) GetCashierById(ctx context.Context, id int64) (entities.Cashier, error) {
	var row entities.Cashier

	query := `
		SELECT
			cashier_id, name, COUNT(cashier_id) OVER() AS row_count
		FROM cashiers WHERE cashier_id = ?
	`
	err := r.mysql.QueryRowxContext(ctx, query, id).StructScan(&row)
	if err != nil {
		logger.Error(ctx, "failed to get cashiers", err, logger.GetCallerTrace())
		if err == sql.ErrNoRows {
			return row, pkg.FormatErrorToCommonError(pkg.ErrCashierNotFound)
		}

		return row, pkg.FormatErrorToCommonError(pkg.ErrFatalQuery)
	}

	return row, nil
}

func (r *repo) CreateCashier(ctx context.Context, req entities.Cashier) (entities.Cashier, error) {
	query := `
		INSERT INTO cashiers (
			name, passcode
		) VALUES (?, ?) 
	`

	row, err := r.mysql.ExecContext(ctx, query, req.Name, req.Passcode)
	if err != nil {
		logger.Error(ctx, "failed to insert cashiers", err, logger.GetCallerTrace())
		return entities.Cashier{}, pkg.FormatErrorToCommonError(pkg.ErrFatalQuery)
	}

	id, _ := row.LastInsertId()

	return entities.Cashier{
		CashierId: id,
		Name:      req.Name,
		Passcode:  req.Passcode,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (r *repo) UpdateCashier(ctx context.Context, req entities.Cashier) error {
	// BEGIN TRANSACTION ----------------------------------------------
	err := dbtx.WithTransaction(ctx, r.mysql, func(tx *sqlx.Tx) error {
		var count int

		query := `
			SELECT count(cashier_id) FROM cashiers WHERE cashier_id = ?
		`
		err := tx.QueryRowxContext(ctx, query, req.CashierId).Scan(&count)
		if err != nil {
			logger.Error(ctx, "failed to get cashiers", err, logger.GetCallerTrace())
			return pkg.ErrFatalQuery
		}

		if count == 0 {
			return pkg.FormatErrorToCommonError(pkg.ErrCashierNotFound)
		}

		query = `
			UPDATE cashiers
				SET name = ?, passcode = ?
			WHERE cashier_id = ?
		`
		_, err = tx.ExecContext(ctx, query, req.Name, req.Passcode, req.CashierId)
		if err != nil {
			logger.Error(ctx, "failed to update cashier", err, logger.GetCallerTrace())
			return pkg.FormatErrorToCommonError(pkg.ErrFatalQuery)
		}

		return err
	})
	// END TRANSACTION -------------------------------------------------

	if err != nil {
		logger.Error(ctx, "failed to commit transaction update cashier", err, logger.GetCallerTrace())
		return err
	}

	return nil
}

func (r *repo) DeleteCashier(ctx context.Context, id int64) error {
	// BEGIN TRANSACTION ----------------------------------------------
	err := dbtx.WithTransaction(ctx, r.mysql, func(tx *sqlx.Tx) error {
		var count int

		query := `
			SELECT count(cashier_id) FROM cashiers WHERE cashier_id = ?
		`
		err := tx.QueryRowxContext(ctx, query, id).Scan(&count)
		if err != nil {
			logger.Error(ctx, "failed to get cashiers", err, logger.GetCallerTrace())
			return pkg.ErrFatalQuery
		}

		if count == 0 {
			return pkg.FormatErrorToCommonError(pkg.ErrCashierNotFound)
		}

		query = `
			DELETE FROM cashiers WHERE cashier_id = ?
		`
		_, err = tx.ExecContext(ctx, query, id)
		if err != nil {
			logger.Error(ctx, "failed to delete cashier", err, logger.GetCallerTrace())
			return pkg.FormatErrorToCommonError(pkg.ErrFatalQuery)
		}

		return err
	})
	// END TRANSACTION -------------------------------------------------

	if err != nil {
		logger.Error(ctx, "failed to commit transaction delete cashier", err, logger.GetCallerTrace())
		return err
	}

	return nil
}
