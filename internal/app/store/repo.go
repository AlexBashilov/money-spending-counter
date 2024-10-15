package store

import (
	"booker/model/apiModels"
	"booker/model/repomodels"
	"context"
)

// ItemsRepository interface
type ItemsRepository interface {
	CreateItems(ctx context.Context, items *repomodels.Items) error
	GetAllItems(ctx context.Context) ([]repomodels.Items, error)
	GetOnlyOneItem(id int) (*apiModels.UserCostItems, error)
	DeleteItems(ctx context.Context, id int) error
	UpdateItems(u *apiModels.UserCostItems, id int) (*apiModels.UserCostItems, error)
	//AddDeletedAt(id int) error
}

// ExpenseRepository interface
type ExpenseRepository interface {
	GetExpenseByDate(period *apiModels.ExpensePeriod) ([]map[string]interface{}, error)
	GetExpenseByItemAndDate(time *apiModels.ExpensePeriod) ([]map[string]interface{}, error)
	GetExpenseSummByPeriodAndItem(time *apiModels.ExpensePeriod) (string, error)
	GetExpenseSummByPeriod(time *apiModels.ExpensePeriod) (string, error)
	AddDeletedTime(int) error
	CheckExist(comparisonSign interface{}) (bool, error)
	GetExpenseSum() ([]map[string]interface{}, error)
	GetExpenseSumByMonth(month int) ([]map[string]interface{}, error)
}
