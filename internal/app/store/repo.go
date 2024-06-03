package store

import "booker/internal/app/model"

type BookerRepository interface {
	CreateItems(items *model.UserCostItems) error
	GetAllItems() (map[string]interface{}, error)
	DeleteItems(id int) error
}
