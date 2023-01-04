package entities

import "time"

type Cashier struct {
	CashierId int64     `json:"cashierId" db:"cashier_id" example:"1"`
	Name      string    `json:"name" db:"name" example:"Kasir 1"`
	Passcode  string    `json:"passcode" db:"passcode" example:"123456"`
	CreatedAt time.Time `json:"createdAt" db:"created_at" example:"2022-04-23T18:25:43.511Z"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at" example:"2022-04-23T18:25:43.511Z"`
	RowCount  int       `json:"-" db:"row_count"`
}

type (
	CashierFilter struct {
		Limit int `json:"limit" query:"limit" validate:"numeric,min=0"`
		Skip  int `json:"skip" query:"skip" validate:"numeric,min=0"`
	}
)
