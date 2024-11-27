package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"github.com/uptrace/bun"

	"booker/model/apiModels"
	"booker/model/repomodels"
)

// ItemsRepo initial items repo
type ExpenseRepo struct {
	client *bun.DB
}

func NewExpenseRepo(client *bun.DB) *ExpenseRepo {
	return &ExpenseRepo{client: client}
}

func (e *ExpenseRepo) GetExpenseByDate(period *apiModels.ExpensePeriod) ([]map[string]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (e *ExpenseRepo) GetExpenseByItemAndDate(time *apiModels.ExpensePeriod) ([]map[string]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (e *ExpenseRepo) GetExpenseSummByPeriodAndItem(time *apiModels.ExpensePeriod) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (e *ExpenseRepo) GetExpenseSummByPeriod(time *apiModels.ExpensePeriod) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (e *ExpenseRepo) AddDeletedTime(i int) error {
	//TODO implement me
	panic("implement me")
}

func (e *ExpenseRepo) CheckExist(comparisonSign interface{}) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (e *ExpenseRepo) AddDeletedAt(id int) error {
	//TODO implement me
	panic("implement me")
}

func (e *ExpenseRepo) GetExpenseSum() ([]map[string]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (e *ExpenseRepo) GetExpenseSumByMonth(month int) ([]map[string]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

// CreateExpense create expense
func (e *ExpenseRepo) CreateExpense(ctx context.Context, expense *repomodels.Expense) error {
	_, err := e.client.NewInsert().
		Model(expense).
		Column("amount").
		Column("date").
		Column("item").
		Exec(ctx)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return errors.Wrap(err, "ошибка внесения суммы в БД")
	}

	var items *repomodels.Items

	if err := e.UpdateItemID(ctx, expense, items); err != nil {
		return err
	}

	return nil
}

// UpdateItemID update item ID
func (e *ExpenseRepo) UpdateItemID(ctx context.Context, expense *repomodels.Expense, items *repomodels.Items) error {
	var result repomodels.Items

	err := e.client.NewSelect().
		Model(items).
		Where("item_name = ?", expense.Item).
		Scan(ctx, &result)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		result = repomodels.Items{}
	}

	err = e.client.NewUpdate().
		Model(expense).
		Where("item = ?", expense.Item).
		Set("item_id = ?", &result.ID).
		Scan(ctx)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("ошибка обновления book_daily_expense.id: %w", err)
	}

	return nil
}

// GetExpenseByItem get expense by item
func (e *ExpenseRepo) GetExpenseByItem(ctx context.Context, itemID int) ([]repomodels.Expense, error) {

	var expense *repomodels.Expense
	var result []repomodels.Expense

	err := e.client.NewSelect().
		Model(expense).
		Column("item").
		Column("amount").
		Column("date").
		Where("item_id = ?", itemID).
		Where("deleted_at is null").
		Scan(ctx, &result)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		result = []repomodels.Expense{}
	}

	return result, nil
}

//// GetExpenseByDate get expense by date
//func (r *BookerRepository) GetExpenseByDate(time *model.ExpensePeriod) ([]map[string]interface{}, error) {
//	rows, err := r.store.db.Query(
//		"SELECT id, amount, date, item FROM book_daily_expense WHERE date >= $1 AND date <= $2 AND deleted_at IS NULL", time.FromDate, time.ToDate)
//	if err != nil {
//		log.Fatal(err)
//	}
//	colNames, err := rows.Columns()
//	if err != nil {
//		log.Fatal(err)
//	}
//	cols := make([]interface{}, len(colNames))
//	colPtrs := make([]interface{}, len(colNames))
//	for i := 0; i < len(colNames); i++ {
//		colPtrs[i] = &cols[i]
//	}
//
//	var querySlice = make([]map[string]interface{}, 0)
//	for rows.Next() {
//		var queryMap = make(map[string]interface{})
//		err = rows.Scan(colPtrs...)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		for i, col := range cols {
//			queryMap[colNames[i]] = col
//		}
//		querySlice = append(querySlice, queryMap)
//	}
//	return querySlice, nil
//}
//
//// GetExpenseByItemAndDate get expense by item and date
//func (r *BookerRepository) GetExpenseByItemAndDate(time *model.ExpensePeriod) ([]map[string]interface{}, error) {
//	rows, err := r.store.db.Query(
//		"SELECT id, amount, date, item FROM book_daily_expense WHERE date >= $1 AND date <= $2 AND item = $3 AND deleted_at IS NULL", time.FromDate, time.ToDate, time.Item)
//	if err != nil {
//		log.Fatal(err)
//	}
//	colNames, err := rows.Columns()
//	if err != nil {
//		log.Fatal(err)
//	}
//	cols := make([]interface{}, len(colNames))
//	colPtrs := make([]interface{}, len(colNames))
//	for i := 0; i < len(colNames); i++ {
//		colPtrs[i] = &cols[i]
//	}
//
//	var querySlice = make([]map[string]interface{}, 0)
//	for rows.Next() {
//		var queryMap = make(map[string]interface{})
//		err = rows.Scan(colPtrs...)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		for i, col := range cols {
//			queryMap[colNames[i]] = col
//		}
//		querySlice = append(querySlice, queryMap)
//	}
//	return querySlice, nil
//}
//
//// GetExpenseSummByPeriodAndItem get summ by period
//func (r *BookerRepository) GetExpenseSummByPeriodAndItem(time *model.ExpensePeriod) (string, error) {
//	var expenseSumm, expenseQuantity float64
//
//	if err := r.store.db.QueryRow(
//		"SELECT SUM(amount) FROM book_daily_expense WHERE date >= $1 AND date <= $2 AND item = $3 AND deleted_at IS NULL",
//		time.FromDate,
//		time.ToDate,
//		time.Item).Scan(&expenseSumm); err != nil {
//		log.Fatal(err)
//	}
//
//	if err := r.store.db.QueryRow(
//		"SELECT count(*) item FROM book_daily_expense WHERE date >= $1 AND date <= $2 AND item = $3 AND deleted_at IS NULL",
//		time.FromDate,
//		time.ToDate,
//		time.Item).Scan(&expenseQuantity); err != nil {
//		log.Fatal(err)
//	}
//
//	formattedResponse := fmt.Sprintf("Вы потратили - %d, Количество трат - %d", int(expenseSumm), int(expenseQuantity))
//	return formattedResponse, nil
//}
//
//// GetExpenseSummByPeriod get expense by period
//func (r *BookerRepository) GetExpenseSummByPeriod(time *model.ExpensePeriod) (string, error) {
//	var expenseSumm, expenseQuantity float64
//
//	if err := r.store.db.QueryRow(
//		"SELECT SUM(amount) FROM book_daily_expense WHERE date >= $1 AND date <= $2 AND deleted_at IS NULL",
//		time.FromDate,
//		time.ToDate).Scan(&expenseSumm); err != nil {
//		log.Fatal(err)
//	}
//
//	if err := r.store.db.QueryRow(
//		"SELECT count(*) item FROM book_daily_expense WHERE date >= $1 AND date <= $2 AND deleted_at IS NULL",
//		time.FromDate,
//		time.ToDate).Scan(&expenseQuantity); err != nil {
//		log.Fatal(err)
//	}
//
//	formattedResponse := fmt.Sprintf("Вы потратили - %d, Количество трат - %d", int(expenseSumm), int(expenseQuantity))
//	return formattedResponse, nil
//}
//
//// AddDeletedTime Add time to DB
//func (r *BookerRepository) AddDeletedTime(itemID int) error {
//	var countedRows int
//
//	if err := r.store.db.QueryRow(
//		"SELECT COUNT(*) FROM book_daily_expense WHERE book_daily_expense.item_id =$1", itemID).
//		Scan(&countedRows); err != nil {
//		return err
//	}
//
//	if countedRows > 0 {
//		sqlStatment := "UPDATE book_daily_expense SET deleted_at =$1 WHERE item_id=$2"
//		_, err := r.store.db.Exec(sqlStatment, time.Now(), itemID)
//		if err != nil {
//			return errors.New("failed to update item_id")
//		}
//	}
//	return nil
//}
//
//// AddDeletedAt Add date to DB
//func (r *BookerRepository) AddDeletedAt(id int) error {
//	_, err := r.store.db.Exec("UPDATE public.book_daily_expense SET deleted_at = $2 WHERE id = $1;", id, time.Now())
//	if err != nil {
//		log.Fatal(err)
//	}
//	return nil
//}
