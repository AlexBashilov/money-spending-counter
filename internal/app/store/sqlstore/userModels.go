package sqlstore

import (
	"booker/internal/app/model"
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
