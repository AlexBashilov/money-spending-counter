package repository

//// GetExpenseSum get expense by sum
//func (r *BookerRepository) GetExpenseSum() ([]map[string]interface{}, error) {
//
//	rows, err := r.store.db.Query(
//		"SELECT item, SUM(amount) FROM book_daily_expense WHERE deleted_at IS NULL GROUP BY item",
//	)
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
//	var mySlice = make([]map[string]interface{}, 0)
//	for rows.Next() {
//		var myMap = make(map[string]interface{})
//		err = rows.Scan(colPtrs...)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		for i, col := range cols {
//			myMap[colNames[i]] = col
//		}
//		mySlice = append(mySlice, myMap)
//	}
//
//	if len(mySlice) < 1 {
//		return nil, errors.New("report is empty")
//	}
//	return mySlice, nil
//}
//
//// GetExpenseSumByMonth get expense by sum and month
//func (r *BookerRepository) GetExpenseSumByMonth(month int) ([]map[string]interface{}, error) {
//
//	rows, err := r.store.db.Query(
//		"SELECT item, SUM(amount) FROM book_daily_expense WHERE deleted_at IS NULL AND EXTRACT('month' from  date) = $1 GROUP BY item", month,
//	)
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
//	var mySlice = make([]map[string]interface{}, 0)
//	for rows.Next() {
//		var myMap = make(map[string]interface{})
//		err = rows.Scan(colPtrs...)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		for i, col := range cols {
//			myMap[colNames[i]] = col
//		}
//		mySlice = append(mySlice, myMap)
//	}
//
//	if len(mySlice) < 1 {
//		return nil, errors.New("report is empty")
//	}
//	return mySlice, nil
//}
