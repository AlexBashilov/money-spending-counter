package sqlstore

import (
	"booker/internal/app/model"
	"errors"
)

func (r *BookerRepository) CreateExpense(u *model.UserExpense) error {
	exists, _ := r.checkItemIsExist(u.Item)
	if exists == false {
		return r.store.db.QueryRow(
			"INSERT INTO book_daily_expense (amount, date, item) VALUES ($1, $2, $3) RETURNING id",
			u.Amount,
			u.Date,
			u.Item,
		).Scan(&u.ID)
	}
	return errors.New("invalid item")
}
