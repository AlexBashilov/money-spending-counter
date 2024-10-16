package store

import (
	"booker/internal/app/model"
)

// BookerRepository interface
type BookerRepository interface {
	CreateItems(items *model.UserCostItems) error
	GetAllItems() ([]map[string]interface{}, error)
	GetOnlyOneItem(id int) (*model.UserCostItems, error)
	DeleteItems(id int) error
	UpdateItems(u *model.UserCostItems, id int) (*model.UserCostItems, error)
	CreateExpense(u *model.UserExpense) error
	GetExpenseByItem(itemID int) ([]map[string]interface{}, error)
	UpdateItemID(item string) error
	GetExpenseByDate(period *model.ExpensePeriod) ([]map[string]interface{}, error)
	GetExpenseByItemAndDate(time *model.ExpensePeriod) ([]map[string]interface{}, error)
	GetExpenseSummByPeriodAndItem(time *model.ExpensePeriod) (string, error)
	GetExpenseSummByPeriod(time *model.ExpensePeriod) (string, error)
	AddDeletedTime(int) error
	CheckExist(comparisonSign interface{}) (bool, error)
	AddDeletedAt(id int) error
	GetExpenseSum() ([]map[string]interface{}, error)
	GetExpenseSumByMonth(month int) ([]map[string]interface{}, error)
}
