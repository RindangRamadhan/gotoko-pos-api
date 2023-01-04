package category

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

func (r *repo) GetCategory(ctx context.Context, req entities.CategoryFilter) ([]entities.Category, response.MetaTpl, error) {
	var (
		pagination response.MetaTpl
		result     struct {
			Rows     []entities.Category
			TotalRow int `json:"total_row"`
		}
	)

	// Build pagination
	skip, limit, offset := request.BuildPagination(req.Skip, req.Limit)

	query := `
		SELECT
			category_id, name, COUNT(category_id) OVER() AS row_count
		FROM categories
		ORDER BY category_id DESC
		LIMIT ? OFFSET ?
	`
	err := r.mysql.SelectContext(ctx, &result.Rows, query, limit, offset)
	if err != nil {
		logger.Error(ctx, "failed to get categories", err, logger.GetCallerTrace())
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

func (r *repo) GetCategoryById(ctx context.Context, id int64) (entities.Category, error) {
	var row entities.Category

	query := `
		SELECT
			category_id, name, COUNT(category_id) OVER() AS row_count
		FROM categories WHERE category_id = ?
	`
	err := r.mysql.QueryRowxContext(ctx, query, id).StructScan(&row)
	if err != nil {
		logger.Error(ctx, "failed to get categories", err, logger.GetCallerTrace())
		if err == sql.ErrNoRows {
			return row, pkg.FormatErrorToCommonError(pkg.ErrCategoryNotFound)
		}

		return row, pkg.FormatErrorToCommonError(pkg.ErrFatalQuery)
	}

	return row, nil
}

func (r *repo) CreateCategory(ctx context.Context, req entities.Category) (entities.Category, error) {
	query := `
		INSERT INTO categories ( name ) 
			VALUES (?) 
	`

	row, err := r.mysql.ExecContext(ctx, query, req.Name)
	if err != nil {
		logger.Error(ctx, "failed to insert categories", err, logger.GetCallerTrace())
		return entities.Category{}, pkg.FormatErrorToCommonError(pkg.ErrFatalQuery)
	}

	id, _ := row.LastInsertId()

	return entities.Category{
		CategoryId: id,
		Name:       req.Name,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}

func (r *repo) UpdateCategory(ctx context.Context, req entities.Category) error {
	// BEGIN TRANSACTION ----------------------------------------------
	err := dbtx.WithTransaction(ctx, r.mysql, func(tx *sqlx.Tx) error {
		var count int

		query := `
			SELECT count(category_id) FROM categories WHERE category_id = ?
		`
		err := tx.QueryRowxContext(ctx, query, req.CategoryId).Scan(&count)
		if err != nil {
			logger.Error(ctx, "failed to get categories", err, logger.GetCallerTrace())
			return pkg.ErrFatalQuery
		}

		if count == 0 {
			return pkg.FormatErrorToCommonError(pkg.ErrCategoryNotFound)
		}

		query = `
			UPDATE categories
				SET name = ?
			WHERE category_id = ?
		`
		_, err = tx.ExecContext(ctx, query, req.Name, req.CategoryId)
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

func (r *repo) DeleteCategory(ctx context.Context, id int64) error {
	// BEGIN TRANSACTION ----------------------------------------------
	err := dbtx.WithTransaction(ctx, r.mysql, func(tx *sqlx.Tx) error {
		var count int

		query := `
			SELECT count(category_id) FROM categories WHERE category_id = ?
		`
		err := tx.QueryRowxContext(ctx, query, id).Scan(&count)
		if err != nil {
			logger.Error(ctx, "failed to get categories", err, logger.GetCallerTrace())
			return pkg.ErrFatalQuery
		}

		if count == 0 {
			return pkg.FormatErrorToCommonError(pkg.ErrCategoryNotFound)
		}

		query = `
			DELETE FROM categories WHERE category_id = ?
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
