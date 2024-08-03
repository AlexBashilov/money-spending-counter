package sqlstore

import (
	"booker/internal/app/model"
	"database/sql"
	"errors"
	"log"
	"time"
)

type BookerRepository struct {
	store *Store
}

func (r *BookerRepository) CreateItems(u *model.UserCostItems) error {
	if err := u.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO book_cost_items (item_name, code, description) VALUES ($1, $2, $3) RETURNING id",
		u.ItemName,
		u.Code,
		u.Description,
	).Scan(&u.ID)
}

func (r *BookerRepository) GetAllItems() ([]map[string]interface{}, error) {

	rows, err := r.store.db.Query(
		"SELECT id, item_name, code, description FROM book_cost_items WHERE deleted_at IS NULL",
	)
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
		return nil, errors.New("No items found")
	}
	return mySlice, nil
}

func (r *BookerRepository) DeleteItems(id int) error {
	_, err := r.store.db.Exec("UPDATE public.book_cost_items SET deleted_at = $2 WHERE id = $1;", id, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (r *BookerRepository) GetOnlyOneItem(itemId int) (*model.UserCostItems, error) {
	var id int
	var itemName string
	var code int
	var description string

	u := &model.UserCostItems{
		ID:          id,
		ItemName:    itemName,
		Code:        code,
		Description: description,
	}
	rows := r.store.db.QueryRow(
		"SELECT id, item_name, code, description FROM book_cost_items WHERE id = $1 AND deleted_at IS NULL",
		itemId,
	).Scan(
		&u.ID,
		&u.ItemName,
		&u.Code,
		&u.Description)
	if errors.Is(rows, sql.ErrNoRows) {
		return nil, nil
	}
	return u, nil
}
