package model

import "github.com/google/uuid"

// CreateItemsRequest struct
type CreateItemsRequest struct {
	ItemName    string    `json:"item_name" validate:"required"`
	GUID        uuid.UUID `json:"guid" validate:"required"`
	Description string    `json:"description" validate:"required"`
}
