package store

import "booker/internal/app/model"

type BookerRepository interface {
	CreateItems(items *model.UserCostItems) error
	GetAllItems() ([]map[string]interface{}, error)
	GetOnlyOneItem(id int) (*model.UserCostItems, error)
	DeleteItems(id int) error
	CreateExpense(u *model.UserExpense) error
}
