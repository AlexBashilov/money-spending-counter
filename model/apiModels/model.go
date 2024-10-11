package apiModels

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"time"
)

// UserCostItems struct items
type UserCostItems struct {
	ID          int       `json:"id"`
	ItemName    string    `json:"item_name"`
	GUID        uuid.UUID `json:"guid"`
	Description string    `json:"description"`
}

// UserExpense struct expense
type UserExpense struct {
	ID     int       `json:"id"`
	Amount float32   `json:"amount"`
	Date   time.Time `json:"date"`
	Item   string    `json:"item"`
	ItemID int       `json:"item_id"`
}

// ExpensePeriod struct
type ExpensePeriod struct {
	FromDate time.Time `json:"fromdate"`
	ToDate   time.Time `json:"todate"`
	Item     string    `json:"item"`
}

// Validate field
func (u *UserCostItems) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.ItemName, validation.Required),
		validation.Field(&u.GUID, validation.Required),
	)
}
