package store

import (
	"booker/model/apiModels"
	"github.com/uptrace/bun"
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

//// CreateExpense create expense
//func (r *BookerRepository) CreateExpense(u *model.UserExpense) error {
//	exists, _ := r.CheckExist(u.Item)
//	if exists {
//		return r.store.db.QueryRow(
//			"INSERT INTO book_daily_expense (amount, date, item) VALUES ($1, $2, $3) RETURNING id",
//			u.Amount,
//			u.Date,
//			u.Item,
//		).Scan(&u.ID)
//	}
//	if err := r.UpdateItemID(u.Item); err != nil {
//		return err
//	}
//	return errors.New("invalid item")
//}
//
//// UpdateItemID update item ID
//func (r *BookerRepository) UpdateItemID(item string) error {
//	var id int
//
//	if err := r.store.db.QueryRow(
//		"SELECT id FROM book_cost_items WHERE item_name = $1",
//		item,
//	).Scan(&id); err != nil {
//		return err
//	}
//
//	sqlStatment := "UPDATE book_daily_expense SET item_id =$1 WHERE item=$2"
//	_, err := r.store.db.Exec(sqlStatment, id, item)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//// GetExpenseByItem get expense by item
//func (r *BookerRepository) GetExpenseByItem(itemID int) ([]map[string]interface{}, error) {
//	rows, err := r.store.db.Query(
//		"SELECT item, amount, date FROM book_daily_expense WHERE item_id = $1 AND deleted_at IS NULL", itemID)
//	if err != nil {
//		log.Fatal(err)
//	}
//
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
