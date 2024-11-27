package store

import (
	"context"

	"booker/model/repomodels"
)

// ItemsRepository interface
type ItemsRepository interface {
	CreateItems(ctx context.Context, items *repomodels.Items) error
	GetAllItems(ctx context.Context) ([]repomodels.Items, error)
	GetOne(ctx context.Context, itemID int) (*repomodels.Items, error)
	DeleteItems(ctx context.Context, id int) error
	UpdateItems(ctx context.Context, u *repomodels.Items, id int) error
	//AddDeletedAt(id int) error
	CheckExist(ctx context.Context, id int) (bool, error)
	CheckItemsDeletedAt(ctx context.Context, id int) (bool, error)
}

// ExpenseRepository interface
type ExpenseRepository interface {
	CreateExpense(ctx context.Context, expense repomodels.Expense) error
	GetExpenseByItem(ctx context.Context, itemID int) ([]repomodels.Expense, error)
	// GetExpenseByDate(period *apiModels.ExpensePeriod) ([]map[string]interface{}, error)
	// GetExpenseByItemAndDate(time *apiModels.ExpensePeriod) ([]map[string]interface{}, error)
	// GetExpenseSummByPeriodAndItem(time *apiModels.ExpensePeriod) (string, error)
	// GetExpenseSummByPeriod(time *apiModels.ExpensePeriod) (string, error)
	// AddDeletedTime(int) error
	// CheckExist(comparisonSign interface{}) (bool, error)
	// GetExpenseSum() ([]map[string]interface{}, error)
	// GetExpenseSumByMonth(month int) ([]map[string]interface{}, error)
}
