package model

import "time"

// CreateExpenseRequest struct
type CreateExpenseRequest struct {
	Amount float32 `json:"amount" validate:"required"`
	Item   string  `json:"item" validate:"required"`
}

// GetExpenseByDateRequest struct
type GetExpenseByDateRequest struct {
	FromDate time.Time `json:"from_date" validate:"required"`
	ToDate   time.Time `json:"to_date" validate:"required"`
}

// ExpenseItemDateRequest struct
type ExpenseItemDateRequest struct {
	FromDate time.Time `json:"from_date" validate:"required"`
	ToDate   time.Time `json:"to_date" validate:"required"`
	Item     string    `json:"item"`
}
