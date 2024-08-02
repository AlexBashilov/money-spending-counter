package store

import (
	"booker/internal/app/model"
)

type BookerRepository interface {
	CreateItems(items *model.UserCostItems) error
	GetAllItems() ([]map[string]interface{}, error)
	GetOnlyOneItem(id int) (*model.UserCostItems, error)
	DeleteItems(id int) error
	CreateExpense(u *model.UserExpense) error
	GetExpenseByItem(itemID int) ([]map[string]interface{}, error)
	UpdateItemID(item string) error
	GeExpenseByDate(period *model.ExpensePeriod) ([]map[string]interface{}, error)
	GeExpenseByItemAndDate(time *model.ExpensePeriod) ([]map[string]interface{}, error)
	GetExpenseSummByPeriodAndItem(time *model.ExpensePeriod) (string, error)
	GetExpenseSummByPeriod(time *model.ExpensePeriod) (string, error)
	AddDeletedTime(int) error
}
