package repomodels

import (
	"time"

	"github.com/uptrace/bun"
)

type Expense struct {
	bun.BaseModel `bun:"table:book_daily_expense"`

	ID        uint32    `json:"id"               bun:"id"`
	Amount    float32   `json:"amount"               bun:"amount"`
	Date      time.Time `json:"date" bun:"date"`
	Item      string    `json:"item"      bun:"item"`
	DeletedAt time.Time `json:"deleted_at" bun:"deleted_at"`
}
