package sqlstore

import (
	"booker/internal/app/model"
	"errors"
	"fmt"
	"log"
	"time"
)

func (r *BookerRepository) CreateExpense(u *model.UserExpense) error {
	exists, _ := r.CheckExist(u.Item)
	if exists {
		return r.store.db.QueryRow(
			"INSERT INTO book_daily_expense (amount, date, item) VALUES ($1, $2, $3) RETURNING id",
			u.Amount,
			u.Date,
			u.Item,
		).Scan(&u.ID)
	}
	if err := r.UpdateItemID(u.Item); err != nil {
		return err
	}
	return errors.New("invalid item")
}

func (r *BookerRepository) UpdateItemID(item string) error {
	var id int

	if err := r.store.db.QueryRow(
		"SELECT id FROM book_cost_items WHERE item_name = $1",
		item,
	).Scan(&id); err != nil {
		return err
	}

	sqlStatment := "UPDATE book_daily_expense SET item_id =$1 WHERE item=$2"
	_, err := r.store.db.Exec(sqlStatment, id, item)
	if err != nil {
		return err
	}
	return nil
}

func (r *BookerRepository) GetExpenseByItem(itemID int) ([]map[string]interface{}, error) {
	rows, err := r.store.db.Query(
		"SELECT item, amount, date FROM book_daily_expense WHERE item_id = $1 AND deleted_at IS NULL", itemID)
	if err != nil {
		log.Fatal(err)
	}

	colNames, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	cols := make([]interface{}, len(colNames))
	colPtrs := make([]interface{}, len(colNames))
	for i := 0; i < len(colNames); i++ {
		colPtrs[i] = &cols[i]
	}

	var querySlice = make([]map[string]interface{}, 0)
	for rows.Next() {
		var queryMap = make(map[string]interface{})
		err = rows.Scan(colPtrs...)
		if err != nil {
			log.Fatal(err)
		}

		for i, col := range cols {
			queryMap[colNames[i]] = col
		}
		querySlice = append(querySlice, queryMap)
	}
	return querySlice, nil
}

func (r *BookerRepository) GeExpenseByDate(time *model.ExpensePeriod) ([]map[string]interface{}, error) {
	rows, err := r.store.db.Query(
		"SELECT id, amount, date, item FROM book_daily_expense WHERE date >= $1 AND date <= $2 AND deleted_at IS NULL", time.FromDate, time.ToDate)
	if err != nil {
		log.Fatal(err)
	}
	colNames, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	cols := make([]interface{}, len(colNames))
	colPtrs := make([]interface{}, len(colNames))
	for i := 0; i < len(colNames); i++ {
		colPtrs[i] = &cols[i]
	}

	var querySlice = make([]map[string]interface{}, 0)
	for rows.Next() {
		var queryMap = make(map[string]interface{})
		err = rows.Scan(colPtrs...)
		if err != nil {
			log.Fatal(err)
		}

		for i, col := range cols {
			queryMap[colNames[i]] = col
		}
		querySlice = append(querySlice, queryMap)
	}
	return querySlice, nil
}

func (r *BookerRepository) GeExpenseByItemAndDate(time *model.ExpensePeriod) ([]map[string]interface{}, error) {
	rows, err := r.store.db.Query(
		"SELECT id, amount, date, item FROM book_daily_expense WHERE date >= $1 AND date <= $2 AND item = $3 AND deleted_at IS NULL", time.FromDate, time.ToDate, time.Item)
	if err != nil {
		log.Fatal(err)
	}
	colNames, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	cols := make([]interface{}, len(colNames))
	colPtrs := make([]interface{}, len(colNames))
	for i := 0; i < len(colNames); i++ {
		colPtrs[i] = &cols[i]
	}

	var querySlice = make([]map[string]interface{}, 0)
	for rows.Next() {
		var queryMap = make(map[string]interface{})
		err = rows.Scan(colPtrs...)
		if err != nil {
			log.Fatal(err)
		}

		for i, col := range cols {
			queryMap[colNames[i]] = col
		}
		querySlice = append(querySlice, queryMap)
	}
	return querySlice, nil
}

func (r *BookerRepository) GetExpenseSummByPeriodAndItem(time *model.ExpensePeriod) (string, error) {
	var expenseSumm, expenseQuantity float64

	if err := r.store.db.QueryRow(
		"SELECT SUM(amount) FROM book_daily_expense WHERE date >= $1 AND date <= $2 AND item = $3 AND deleted_at IS NULL",
		time.FromDate,
		time.ToDate,
		time.Item).Scan(&expenseSumm); err != nil {
		log.Fatal(err)
	}

	if err := r.store.db.QueryRow(
		"SELECT count(*) item FROM book_daily_expense WHERE date >= $1 AND date <= $2 AND item = $3 AND deleted_at IS NULL",
		time.FromDate,
		time.ToDate,
		time.Item).Scan(&expenseQuantity); err != nil {
		log.Fatal(err)
	}

	formattedResponse := fmt.Sprintf("Вы потратили - %d, Количество трат - %d", int(expenseSumm), int(expenseQuantity))
	return formattedResponse, nil
}

func (r *BookerRepository) GetExpenseSummByPeriod(time *model.ExpensePeriod) (string, error) {
	var expenseSumm, expenseQuantity float64

	if err := r.store.db.QueryRow(
		"SELECT SUM(amount) FROM book_daily_expense WHERE date >= $1 AND date <= $2 AND deleted_at IS NULL",
		time.FromDate,
		time.ToDate).Scan(&expenseSumm); err != nil {
		log.Fatal(err)
	}

	if err := r.store.db.QueryRow(
		"SELECT count(*) item FROM book_daily_expense WHERE date >= $1 AND date <= $2 AND deleted_at IS NULL",
		time.FromDate,
		time.ToDate).Scan(&expenseQuantity); err != nil {
		log.Fatal(err)
	}

	formattedResponse := fmt.Sprintf("Вы потратили - %d, Количество трат - %d", int(expenseSumm), int(expenseQuantity))
	return formattedResponse, nil
}

func (r *BookerRepository) AddDeletedTime(itemId int) error {
	var countedRows int

	if err := r.store.db.QueryRow(
		"SELECT COUNT(*) FROM book_daily_expense WHERE book_daily_expense.item_id =$1", itemId).
		Scan(&countedRows); err != nil {
		return err
	}

	if countedRows > 0 {
		sqlStatment := "UPDATE book_daily_expense SET deleted_at =$1 WHERE item_id=$2"
		_, err := r.store.db.Exec(sqlStatment, time.Now(), itemId)
		if err != nil {
			return errors.New("failed to update item_id")
		}
	}
	return nil
}

func (r *BookerRepository) AddDeletedAt(id int) error {
	_, err := r.store.db.Exec("UPDATE public.book_daily_expense SET deleted_at = $2 WHERE id = $1;", id, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
