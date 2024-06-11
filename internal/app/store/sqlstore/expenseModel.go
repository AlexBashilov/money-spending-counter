package sqlstore

import (
	"booker/internal/app/model"
	"errors"
	"fmt"
	"log"
)

func (r *BookerRepository) CreateExpense(u *model.UserExpense) error {
	exists, _ := r.checkItemIsExist(u.Item)
	if exists == true {
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
		errors.New("failed to find item in book_cost_items")
	}

	sqlStatment := "UPDATE book_daily_expense SET item_id =$1 WHERE item=$2"
	_, err := r.store.db.Exec(sqlStatment, id, item)
	if err != nil {
		return errors.New("failed to update item_id")
	}
	fmt.Println("ITEM - ", id)
	return nil
}

func (r *BookerRepository) GetExpenseByItem(itemID int) ([]map[string]interface{}, error) {
	rows, err := r.store.db.Query(
		"SELECT item, amount, date FROM book_daily_expense WHERE item_id = $1", itemID)
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
		"SELECT id, amount, date, item FROM book_daily_expense WHERE date >= $1 AND date <= $2", time.FromDate, time.ToDate)
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
		"SELECT id, amount, date, item FROM book_daily_expense WHERE date >= $1 AND date <= $2 AND item = $3", time.FromDate, time.ToDate, time.Item)
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
