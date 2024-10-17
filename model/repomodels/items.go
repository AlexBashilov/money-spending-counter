package repomodels

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type Items struct {
	bun.BaseModel `bun:"table:book_cost_items"`

	ID          uint32    `json:"id"               bun:"id"`
	ItemName    string    `json:"item_name"      bun:"item_name"`
	GUID        uuid.UUID `json:"guid"    bun:"guid"`
	Description string    `json:"description"   bun:"description"`
	DeletedAt   time.Time `json:"deleted_at" bun:"deleted_at"`
}
