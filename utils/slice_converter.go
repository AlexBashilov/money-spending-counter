package utils

import (
	"database/sql"
	"errors"
	"log"
)

func ConverToSlice(rows sql.Rows) ([]map[string]interface{}, error) {
	colNames, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	cols := make([]interface{}, len(colNames))
	colPtrs := make([]interface{}, len(colNames))
	for i := 0; i < len(colNames); i++ {
		colPtrs[i] = &cols[i]
	}

	var mySlice = make([]map[string]interface{}, 0)
	for rows.Next() {
		var myMap = make(map[string]interface{})
		err = rows.Scan(colPtrs...)
		if err != nil {
			log.Fatal(err)
		}

		for i, col := range cols {
			myMap[colNames[i]] = col
		}
		mySlice = append(mySlice, myMap)
	}

	if len(mySlice) < 1 {
		return nil, errors.New("no items found")
	}
	return mySlice, nil
}
