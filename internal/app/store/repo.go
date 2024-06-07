package store

import (
	"booker/internal/app/model"
	"time"
)

type BookerRepository interface {
	CreateItems(items *model.UserCostItems) error
	GetAllItems() ([]map[string]interface{}, error)
	GetOnlyOneItem(id int) (*model.UserCostItems, error)
	DeleteItems(id int) error
	CreateExpense(u *model.UserExpense) error
	GetExpenseByItem(itemID int) ([]map[string]interface{}, error)
	UpdateItemID(item string) error
	GeExpenseByDate(time time.Time) ([]map[string]interface{}, error)
}
