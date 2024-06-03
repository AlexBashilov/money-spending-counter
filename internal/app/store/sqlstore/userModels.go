package sqlstore

import (
	"booker/internal/app/model"
	"fmt"
	"log"
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

func (r *BookerRepository) GetAllItems() (map[string]interface{}, error) {
	rows, err := r.store.db.Query(
		"SELECT id, item_name, code, description FROM book_cost_items",
	)
	if err != nil {
		log.Fatal(err)
	}

	cols, _ := rows.Columns()
	m := make(map[string]interface{})

	for rows.Next() {

		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}

	}
	fmt.Println(m)
	return m, nil
}

func (r *BookerRepository) DeleteItems(id int) error {
	_, err := r.store.db.Exec("DELETE FROM  public.book_cost_items WHERE id = $1;", id)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
