package entities

import "time"

type Category struct {
	CategoryId int64     `json:"categoryId" db:"category_id" example:"1"`
	Name       string    `json:"name" db:"name" example:"Kasir 1"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at" example:"2022-04-23T18:25:43.511Z"`
	UpdatedAt  time.Time `json:"updatedAt" db:"updated_at" example:"2022-04-23T18:25:43.511Z"`
	RowCount   int       `json:"-" db:"row_count"`
}

type (
	CategoryFilter struct {
		Limit int `json:"limit" query:"limit" validate:"numeric,min=0"`
		Skip  int `json:"skip" query:"skip" validate:"numeric,min=0"`
	}
)
