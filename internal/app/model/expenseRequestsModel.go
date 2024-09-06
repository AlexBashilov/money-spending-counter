package model

type CreateExpenseRequest struct {
	Amount float32 `json:"amount" validate:"required"`
	Item   string  `json:"item" validate:"required"`
}
