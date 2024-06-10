package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

// Структура и описание статей затрат
type UserCostItems struct {
	ID          int    `json:"id"`
	ItemName    string `json:"item_name"`
	Code        int    `json:"code"`
	Description string `json:"description"`
}

// Структура дневных затрат
type UserExpense struct {
	ID     int       `json:"id"`
	Amount float32   `json:"amount"`
	Date   time.Time `json:"date"`
	Item   string    `json:"item"`
	ItemId int       `json:"item_id"`
}

type ExpensePeriod struct {
	FromDate time.Time `json:"fromdate"`
	ToDate   time.Time `json:"todate"`
	Item     string    `json:"item"`
}

// /Validate ...
func (u *UserCostItems) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.ItemName, validation.Required),
		validation.Field(&u.Code, validation.Required),
	)
}
